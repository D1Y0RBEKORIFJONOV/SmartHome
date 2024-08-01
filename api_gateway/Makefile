# Makefile

# Переменные
SWAGGER_CMD = swag
GO_RUN_CMD = go run

# Команда для инициализации Swagger
swagger-init:
	$(SWAGGER_CMD) init -g internal/handler/handler.go

# Команда для запуска приложения
run:
	$(GO_RUN_CMD) cmd/app/main.go

# Команда по умолчанию, которая сначала инициализирует Swagger, а затем запускает приложение
all: swagger-init run
