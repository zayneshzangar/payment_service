proto:
	protoc --go_out=. --go-grpc_out=. proto/payment.proto

run:
	go run cmd/main.go

goGet:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go unstall google.golang.org/grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/lib/pq
	go get -u github.com/lib/pq
	go install github.com/confluentinc/confluent-kafka-go/kafka
	go get -u github.com/confluentinc/confluent-kafka-go/kafka
