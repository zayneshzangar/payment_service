package rest

import (
	"payment_service/internal/middleware"

	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(paymentHandler *PaymentHandler) *mux.Router {
	router := mux.NewRouter()

	// Оборачиваем handler.Payment в middleware
	router.Handle("/payment/{paymentID}", middleware.JWTMiddleware(http.HandlerFunc(paymentHandler.Payment))).Methods("POST")

	return router
}
