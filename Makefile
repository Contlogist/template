include .env.example
export

TAG?=latest
NAME:=catalog_m
DOCKER_REPOSITORY:=telepro
HARBOR_ADDRESS:=harbor.legchelife.ru:8443
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=$(shell grep 'VERSION' cmd/app/version.go | awk '{ print $$4 }' | tr -d '"')

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

local-latest: ### build local latest
	swag init -g internal/controller/http/v1/router.go
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_IMAGE_NAME):latest
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_IMAGE_NAME):$(TAG)

harbor-latest:
	swag init -g internal/controller/http/v1/router.go
	docker buildx build --platform linux/amd64  -t $(DOCKER_IMAGE_NAME):$(VERSION) -f ./building/latest/Dockerfile .
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) $(HARBOR_ADDRESS)/$(DOCKER_IMAGE_NAME):$(VERSION)
	docker push $(HARBOR_ADDRESS)/$(DOCKER_IMAGE_NAME):$(VERSION)
	docker buildx build --platform linux/amd64  -t $(DOCKER_IMAGE_NAME):latest -f ./building/latest/Dockerfile .
	docker tag $(DOCKER_IMAGE_NAME):latest $(HARBOR_ADDRESS)/$(DOCKER_IMAGE_NAME):latest
	docker push $(HARBOR_ADDRESS)/$(DOCKER_IMAGE_NAME):latest

harbor-latest-debug:
	swag init -g internal/controller/http/v1/router.go
	docker buildx build --platform linux/amd64 -t $(DOCKER_IMAGE_NAME):latest.debug -f ./building/debug/Dockerfile .
	docker tag $(DOCKER_IMAGE_NAME):latest.debug $(HARBOR_ADDRESS)/$(DOCKER_IMAGE_NAME):latest.debug
	docker push $(HARBOR_ADDRESS)/$(DOCKER_IMAGE_NAME):latest.debug

compose-up: ### Run docker-compose
	docker-compose up --build
.PHONY: compose-up

compose-up-integration-test: ### Run docker-compose with integration test
	docker-compose up --build --abort-on-container-exit --exit-code-from integration
.PHONY: compose-up-integration-test

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down

swag-v1: ### swag init
	swag init -g internal/controller/http/v1/router.go
.PHONY: swag-v1

run: swag-v1 ### swag run
	swag init -g internal/controller/http/v1/router.go
	go mod tidy && go mod download && \
	go build git.legchelife.ru/gitlab-instance-7d441567/catalog_m/cmd/app
	DISABLE_SWAGGER_HTTP_HANDLER='' GIN_MODE=debug CGO_ENABLED=0 go run -tags migrate ./cmd/app
.PHONY: run

docker-rm-volume: ### remove docker volume
	docker volume rm go-clean-template_pg-data
.PHONY: docker-rm-volume

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

linter-hadolint: ### check by hadolint linter
	git ls-files --exclude='Dockerfile*' --ignored | xargs hadolint
.PHONY: linter-hadolint

linter-dotenv: ### check by dotenv linter
	dotenv-linter
.PHONY: linter-dotenv

test: ### run test
	go test -v -cover -race ./internal/...
.PHONY: test

integration-test: ### run integration-test
	go clean -testcache && go test -v ./integration-test/...
.PHONY: integration-test

mock: ### run mockery
	mockery --all -r --case snake
.PHONY: mock

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'migrate_name'
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path migrations -database '$(PG_URL)?sslmode=disable' up
.PHONY: migrate-up
