.PHONY: default run run-with-docs build test docs clean

# variables
APP_NAME=gopportunities
MAIN_PATH=cmd/api/main.go

# tasks
default: run-with-docs

run:
	@go run $(MAIN_PATH)
run-with-docs:
	@swag init -g $(MAIN_PATH)
	@go run $(MAIN_PATH)
build:
	@go build -o (APP_NAME) $(MAIN_PATH)
test:
	@go test ./ ...
docs:
	@swag init -g $(MAIN_PATH)
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs