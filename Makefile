run:
	go run main.go

migrate:
	go run cmd/migrate/main.go

swag:
	swag init

proto:
	buf lint
	buf generate
