app: 
	make doc
	make run

run:
	go run cmd/main/main.go

doc:
	swag init -g ./internal/handlers/manager/manager.go