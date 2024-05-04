.PHONY: start

start:
	DATA_SOURCE_URL="root:verysecretpass@tcp(127.0.0.1:3306)/order" \
	APPLICATION_PORT=3000 \
	ENV=development \
	PAYMENT_SERVICE_URL=localhost:3001 \
	go run cmd/main.go
