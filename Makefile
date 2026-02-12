DOCKER ?= docker

GO_TEST_EXTRA_ARGS ?=

INPUT_OPENAPI_SCHEMA ?=
OPENAPI_PATCHED_SCHEMA ?= openapi-patched.json
OPENAPI_SCHEMA ?= openapi.json

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

# XXX(ndhoule):
# These changes to the OpenAPI schema are super unfortunate. In many (but an unknown number of)
# cases, these are actual fixes to issues in our OpenAPI schema. In some places, the fix is not
# how we'd make the change upstream (e.g. PartialSite).
#
# This mega-gross chain of stuff:
#
# - Fixes some egregious errors in the OpenAPI schema (e.g. "type": "datetime")
# - Recursively merges the contents of openapi.json.patch into the external OpenAPI schema to add
#   missing or internal properties, add deprecated APIs back into the schema for backward
#   compatibility, etc.
#  - Then makes a chain of fixes to the merged schema
#
# This is order sensitive: E.g. we _must_ merge the patch in before we mutate the schema.
$(OPENAPI_SCHEMA): $(INPUT_OPENAPI_SCHEMA) openapi.json.patch ## Generate a patched copy of Netlify's external OpenAPI schema suitable for use by the API client generator.
ifndef INPUT_OPENAPI_SCHEMA
	@echo "INPUT_OPENAPI_SCHEMA must be set to a local path"
	@exit 1
endif
	@jq -s 'reduce .[] as $$obj ({}; . * $$obj)' $^ \
		| sed 's|"type": "status"|"type": "string"|g' \
		| sed 's|"type": "hash"|"type": "object"|g' \
		| sed 's|"type": "datetime"|"type": "string", "format": "date-time"|g' \
		| jq 'del(.components.schemas.Account.properties.billing_details)' \
		| jq '.components.schemas.Account.required |= map(select(. != "billing_details"))' \
		| jq 'del(.components.schemas.Repo.required)' \
		| jq '.components.schemas.DnsZone.required |= map(select(. != "domain"))' \
		| jq '.paths["/api/v1/accounts/{account_id}/env/{env_key}"].put.parameters |= map(select(.name != "values"))' \
		| jq '.paths["/api/v1/accounts/{account_id}/env/{env_key}"].put.parameters |= map(select(.name != "scopes"))' \
		| jq '.paths["/api/v1/accounts/{account_id}/env/{env_key}"].put.parameters |= map(select(.name != "is_secret"))' \
		| jq '.components.schemas.EnvVar.required |= map(select(. != "scopes"))' \
		| jq 'del(.components.schemas.LogDrainServiceConfig.required)' \
		| jq '.components.schemas.Site.required |= map(select(. != "password"))' \
		| jq '.components.schemas.Site.required |= map(select(. != "password_context"))' \
		| jq '.components.schemas.Site.required |= map(select(. != "password_hash"))' \
		| jq '.components.schemas.PartialSite = .components.schemas.Site' \
		| jq 'del(.components.schemas.PartialSite.required)' \
		| jq 'del(.components.schemas.DnsRecord.properties.target)' \
		| jq '.components.schemas.DnsRecord.required |= map(select(. != "target"))' \
		> $@

clean-client: ## Remove existing generated API client files.
	@rm -rf internal/netlifyapi
.PHONY: clean-client

client: clean-client ## Generate an API client from the OpenAPI schema.
	@$(DOCKER) run --rm \
		--mount type=bind,src=$(PWD),dst=/local,rw,z \
		docker.io/openapitools/openapi-generator-cli \
		generate \
		--additional-properties=disallowAdditionalPropertiesIfNotPresent=false,isGoSubmodule=true,packageName=netlifyapi,withGoMod=false \
		--global-property apiDocs=false,modelDocs=false,apiTests=false,modelTests=false \
		-i /local/openapi.json \
		-g go \
		-o /local/internal/netlifyapi ; \
	sed -i 's/int32/int64/g' internal/netlifyapi/model_*.go ; \
	sed -i 's/int32/int64/g' internal/netlifyapi/api_*.go ; \
	sed -i 's/return e.error/return fmt.Sprintf("%s: %s", e.error, e.body)/g' internal/netlifyapi/client.go
.PHONY: client

test: ## Test the go code.
	go test -v ./...
.PHONY: test

testacc: ## Test the go code and run acceptance tests.
	TF_ACC=1 go test ./... -v $(GO_TEST_EXTRA_ARGS)
.PHONY: testacc
