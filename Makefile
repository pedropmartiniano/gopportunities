.PHONY: default run build test docs clean

APP_NAME=gopportunities

default: run

run:
	@go run .

build:
	@go build -o $(APP_NAME) main.go

docs:
	@swag init

clean:
	@rm -f $(APP_NAME)
	rm -rf ./docs