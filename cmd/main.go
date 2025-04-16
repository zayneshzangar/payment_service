package main

import (
	"log"
	"net/http"
	grpc "payment_service/internal/delivery/gprc"
	"payment_service/internal/delivery/rest"
	"payment_service/internal/repository"
	"payment_service/internal/service/grpc_service"
	"payment_service/internal/service/payment_service"
)

var address = []string{"localhost:9091", "localhost:9092", "localhost:9093"}

func main() {
	repo, err := repository.NewDatabaseConnection(repository.Postgres)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer repo.Close()

	paymentService, err := payment_service.NewPaymentService(repo, address)
	if err != nil {
		log.Fatalf("NewPaymentService: %v", err)
	}
	paymentHandler := rest.NewPaymentHandler(paymentService)
	router := rest.NewRouter(paymentHandler)

	grpcService := grpc_service.NewGrpcService(repo)
	go grpc.StartGRPCServer(grpcService)

	port := ":8082"
	log.Println("Starting REST API server on", port)
	log.Fatal(http.ListenAndServe(port, corsMiddleware(router)))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true") // 🔥 ВАЖНО

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
