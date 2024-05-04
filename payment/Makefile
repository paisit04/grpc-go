.PHONY: start

start:
	DATA_SOURCE_URL="root:verysecretpass@tcp(127.0.0.1:3306)/payment" \
	APPLICATION_PORT=3001 \
	ENV=development \
	go run cmd/main.go
