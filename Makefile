SWAGGER_SPEC := swagger.yml

help: ## Show this help.
		@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.PHONY: help

all: deps openapi_generate generate test testacc # build
.PHONY: all

build: ## Build Terraform provider.
	go install .
.PHONY: build

deps: ## Install dependencies.
	go mod download
.PHONY: deps

generate: ## Generate Terraform docs.
	go generate
.PHONY: generate

openapi_generate: ## Generate the go code from the OpenAPI spec.
	docker run --rm \
		-v ${PWD}:/local openapitools/openapi-generator-cli generate \
		--additional-properties=disallowAdditionalPropertiesIfNotPresent=false,isGoSubmodule=true,packageName=netlifyapi,withGoMod=false \
		--global-property apiDocs=false,modelDocs=false,apiTests=false,modelTests=false \
		-i /local/openapi.json \
		-g go \
		-o /local/internal/netlifyapi ; \
	sed -i '' 's/int32/int64/g' internal/netlifyapi/model_*.go ; \
	sed -i '' 's/int32/int64/g' internal/netlifyapi/api_*.go ; \
	sed -i '' 's/return e.error/return fmt.Sprintf("%s: %s", e.error, e.body)/g' internal/netlifyapi/client.go
.PHONY: openapi_generate

test: ## Test the go code.
	go test -v ./...
.PHONY: test

testacc: ## Test the go code and run acceptance tests.
	TF_ACC=1 go test ./... -v $(TESTARGS)
# -timeout 120m
.PHONY: testacc
