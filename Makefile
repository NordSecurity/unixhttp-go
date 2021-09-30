.PHONY: dep fmt qa lint vet test coverage help

dep: ## Install app dependencies
	go mod tidy
	go mod vendor

fmt: ## Reformat the code
	go fmt ./...

qa: lint vet test ## Check code quality

lint: ## Lint the code
	test -z $$(gofmt -l . | grep -v vendor/) || (echo "Formatting issues found in:" $$(gofmt -l . | grep -v vendor/) && exit 1)

vet: ## Vet the code
	go vet ./...

test: ## Run the tests
	go test ./... -v -coverprofile .testCoverage.txt

coverage: test ## Show test coverage info in the browser
	go tool cover -html .testCoverage.txt

help: ## Show available commands
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | awk 'BEGIN {FS = ":.*?## "}; \
	{printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
