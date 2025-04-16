package entity

type Payment struct {
	ID      int64   `json:"id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
	OrderID int64   `json:"order_id"`
	UserID  int64   `json:"user_id"`
}

type PaymentRequest struct {
	CardNum string `json:"card_number"`
	CVV     string `json:"cvv"`
	ExpDate string `json:"exp_date"`
	Name    string `json:"name"`
}

type OrderNotification struct {
    OrderID int64  `json:"order_id"`
    Status  string `json:"status"`
}

type PaymentNotification struct {
    OrderID int64 `json:"order_id"`
    Status string `json:"status"`
}
