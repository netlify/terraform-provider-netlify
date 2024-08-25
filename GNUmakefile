.PHONY: all build deps generate help test openapi_generate
SWAGGER_SPEC=swagger.yml

help: ## Show this help.
		@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

all: deps openapi_generate generate test testacc # build

build: ## Build Terraform provider.
	go install .

deps: ## Install dependencies.
	go mod download

generate: ## Generate Terraform docs.
	go generate

openapi_generate: ## Generate the go code from the OpenAPI spec.
	docker run --rm \
		-v ${PWD}:/local openapitools/openapi-generator-cli generate \
		--additional-properties=disallowAdditionalPropertiesIfNotPresent=false,isGoSubmodule=true,packageName=netlifyapi,withGoMod=false \
		--global-property apiDocs=false,modelDocs=false,apiTests=false,modelTests=false \
		-i /local/openapi.json \
		-g go \
		-o /local/internal/netlifyapi ; \
	sed -i '' 's/int32/int64/g' internal/netlifyapi/model_*.go ; \
	sed -i '' 's/int32/int64/g' internal/netlifyapi/api_*.go

test: ## Test the go code.
	go test -v ./...

testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS)
# -timeout 120m
