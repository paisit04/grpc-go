.PHONY: start-db start-server

start-db:
	docker run -d -p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=verysecretpass \
	-e MYSQL_DATABASE=order mysql

start-server:
	DATA_SOURCE_URL="root:verysecretpass@tcp(127.0.0.1:3306)/order" \
	APPLICATION_PORT=3000 \
	ENV=development \
	go run cmd/main.go
