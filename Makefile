PROJECT_NAME := kabikha
PKG_LIST := $(shell go list ${PROJECT_NAME}/... | grep -v /vendor/)

.PHONY: run test coverage clean

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


########################
### DEVELOP and TEST ###
########################

run:
	# booting up dependecy containers
	@docker-compose up -d db

	# booting up application container in foreground to trail the logs
	@docker-compose up --build server


test: ## Run unittests
	@go test -cover -short ${PKG_LIST}

coverage: ## Generate global code coverage report
	@go tool cover -func=cov.out

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)
	@docker-compose down
