help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Builds a bin
	@go mod tidy
	@go build -o ./bin/chip8emu ./main.go

run: ## Runs the built bin
	@./bin/chip8emu

format: ## Formats the the code
	go fmt .
	go fmt chip8sys/*


.PHONY: help build run format
