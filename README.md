go get github.com/golang-jwt/jwt/v5
go get go.uber.org/mock/gomock
go get github.com/stretchr/testify/assert


# Для PaymentRepository
mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/mocks.go -package=mocks

# Для KafkaProducer
mockgen -source=internal/kafka/interfaces.go -destination=internal/kafka/mocks/mocks.go -package=mocks

go clean -testcache
go test -v ./internal/service/payment_service

go test -cover ./internal/service/payment_service

