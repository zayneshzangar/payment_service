package repository

import (
	"context"
	"database/sql"
	"fmt"

	"payment_service/internal/entity"
)

type PostgresPaymentRepository struct {
	db *sql.DB
}

func (r *PostgresPaymentRepository) CreatePayment(ctx context.Context, payment *entity.Payment) (*entity.Payment, error) {
	query := `INSERT INTO payments (user_id, order_id, amount, status) VALUES ($1, $2, $3, 'pending') RETURNING id`
	err := r.db.QueryRowContext(ctx, query, payment.UserID, payment.OrderID, payment.Amount).Scan(&payment.ID)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PostgresPaymentRepository) GetPayment(ctx context.Context, id int64) (*entity.Payment, error) {
	query := `SELECT id, amount, order_id, status FROM payments WHERE id = $1`
	var payment entity.Payment
	err := r.db.QueryRowContext(ctx, query, id).Scan(&payment.ID, &payment.Amount, &payment.OrderID, &payment.Status)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *PostgresPaymentRepository) UpdatePayment(ctx context.Context, payment *entity.Payment) error {
	query := `UPDATE payments SET status = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, payment.Status, payment.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresPaymentRepository) DeletePayment(ctx context.Context, id string) error {
	query := `DELETE FROM payments WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresPaymentRepository) Close() error {
	return r.db.Close()
}

func (r *PostgresPaymentRepository) ValidateCard(ctx context.Context, cardNum, cvv, expDate, name string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM cards WHERE number = $1 AND cvv = $2 AND exp_date = $3 AND name = $4)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, cardNum, cvv, expDate, name).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *PostgresPaymentRepository) CheckAndDeductBalance(ctx context.Context, cardNum string, amount float64) (bool, error) {
	// Начинаем транзакцию
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return false, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback() // откатим транзакцию в случае ошибки

	var balance float64

	// Проверяем баланс с блокировкой строки (FOR UPDATE)
	query := `SELECT balance FROM cards WHERE number = $1 FOR UPDATE`
	err = tx.QueryRowContext(ctx, query, cardNum).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("card number not found")
		}
		return false, err
	}
	// Проверяем достаточность средств
	if balance < amount {
		return false, nil
	}

	// Списываем деньги с карты
	updateQuery := `UPDATE cards SET balance = balance - $1 WHERE number = $2`
	_, err = tx.ExecContext(ctx, updateQuery, amount, cardNum)
	if err != nil {
		return false, fmt.Errorf("failed to update balance: %w", err)
	}

	// Подтверждаем транзакцию
	if err = tx.Commit(); err != nil {
		return false, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return true, nil
}
