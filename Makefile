.PHONY: swagger
swagger:
	go run github.com/swaggo/swag/v2/cmd/swag init -g main.go