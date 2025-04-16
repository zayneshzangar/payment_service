package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"payment_service/internal/entity"
	"payment_service/internal/service"

	"github.com/gorilla/mux"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (s *PaymentHandler) Payment(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Payment")
	// Получаем paymentID из URL
	vars := mux.Vars(r)
	paymentID, err := strconv.ParseInt(vars["paymentID"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid payment ID",
		})
		return
	}

	var req entity.PaymentRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	err = s.service.PaymentProcess(context.Background(), paymentID, req)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Payment processing failed",
		})
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Payment successful",
	})
}
