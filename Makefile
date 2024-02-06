include .env
export

YQ := yq
LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)

TAG?=latest
NAME:=template
DOCKER_REPOSITORY:=telepro
HARBOR_ADDRESS:=harbor.legchelife.ru:8443
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
export VERSION := $(shell grep '^APP_VERSION=' .env | awk -F= '{print $$2}')



harbor-latest:
	go generate ./ent
	swag init -g internal/controller/http/v1/router.go
	@sed -i '' 's/"version": "\(.*\)"/"version": "$(VERSION)"/' ./docs/swagger.json
	docker buildx build --platform linux/amd64  -t $(DOCKER_IMAGE_NAME):$(VERSION) -f ./building/latest/Dockerfile .
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) $(HARBOR_ADDRESS)/$(DOCKER_IMAGE_NAME):$(VERSION)
	docker push $(HARBOR_ADDRESS)/$(DOCKER_IMAGE_NAME):$(VERSION)
	docker buildx build --platform linux/amd64  -t $(DOCKER_IMAGE_NAME):latest -f ./building/latest/Dockerfile .
	docker tag $(DOCKER_IMAGE_NAME):latest $(HARBOR_ADDRESS)/$(DOCKER_IMAGE_NAME):latest
	docker push $(HARBOR_ADDRESS)/$(DOCKER_IMAGE_NAME):latest

harbor-latest-debug:
	go generate ./ent
	swag init -g internal/controller/http/v1/router.go
	docker buildx build --platform linux/amd64 -t $(DOCKER_IMAGE_NAME):latest.debug -f ./building/debug/Dockerfile .
	docker tag $(DOCKER_IMAGE_NAME):latest.debug $(HARBOR_ADDRESS)/$(DOCKER_IMAGE_NAME):latest.debug
	docker push $(HARBOR_ADDRESS)/$(DOCKER_IMAGE_NAME):latest.debug

compose-up: ### Run docker-compose
	swag init -g internal/controller/http/v1/router.go
	go mod tidy && go mod download && \
	docker-compose up --build
.PHONY: compose-up

compose-up-integration-test: ### Run docker-compose with integration test
	docker-compose up --build --abort-on-container-exit --exit-code-from integration
.PHONY: compose-up-integration-test

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down


run: run ### swag run
	#go generate ./ent
	swag init -g internal/controller/http/v1/router.go
	go mod tidy && go mod download && \
	go build git.legchelife.ru/root/template/cmd/app
	DISABLE_SWAGGER_HTTP_HANDLER='' GIN_MODE=debug CGO_ENABLED=0 go run -tags migrate ./cmd/app
.PHONY: run

ent-gen: ### run ent generate
	go generate ./ent
.PHONY: ent-gen

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

migrate-create:  ### create new migration NAME=...
	migrate create -ext sql -dir migrations/$(VERSION) $(NAME)
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path migrations/$(VERSION) -database '$(PG_URL)?sslmode=disable' up
.PHONY: migrate-up

migrate-force: ### migration force
	migrate -path migrations/$(VERSION) -database '$(PG_URL)?sslmode=disable' force $(FORCE)
.PHONY: migrate-force

migrate-down: ### migration down
	migrate -path migrations/$(VERSION) -database '$(PG_URL)?sslmode=disable' down
.PHONY: migrate-down

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out
