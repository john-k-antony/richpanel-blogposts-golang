export GOBIN = $(shell pwd)/toolbin

GOIMPORTS = $(GOBIN)/goimports
GOLINT = $(GOBIN)/golint
GOLINTCI = $(GOBIN)/golangci-lint
STATICCHECK = $(GOBIN)/staticcheck
OAPI = $(GOBIN)/oapi-codegen
MOCKERY = $(GOBIN)/mockery
SWAGGER = $(GOBIN)/swagger

.PHONY: build
build:
	@mkdir -p build
	@go build -o build/app

.PHONY: run
run: build	
	@./build/app

.PHONY: genswag
genswag:
	@$(SWAGGER) validate --skip-warnings swagger/BlogPost2.0.yaml
	@rm -rf models/apimodels restapi/operations
	@$(SWAGGER) flatten --format=yaml -o=swagger/swagger-flat.tmp.yaml swagger/BlogPost2.0.yaml
	@$(SWAGGER) generate server \
		-f swagger/swagger-flat.tmp.yaml \
		-C swagger/config.yaml \
		-A Richpanel-BlogPost \
		-P models.Principal \
		-m models/apimodels \
		--struct-tags "json" \
		--struct-tags "bson"

.PHONY: genoapi
genoapi:
	@$(OAPI) --config=oapi/config.yaml oapi/BlogPost.yaml

.PHONY: tools
tools:
	@rm -rf toolbin
	@go install golang.org/x/lint/golint
	@go install honnef.co/go/tools/cmd/staticcheck
	@go install github.com/go-swagger/go-swagger/cmd/swagger
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen
	@go install golang.org/x/tools/cmd/goimports
	@go install github.com/vektra/mockery/v2
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: fmt
fmt:
	@$(GOIMPORTS) -l -w .

.PHONY: lint
lint:
	@./ci/lint.sh

.PHONY: genmock
genmock:
	@$(MOCKERY) --all --keeptree 

.PHONY: test
test:
	@go test -v -cover -coverprofile=c.out ./... 
	@go tool cover -html=c.out -o coverage.html

.PHONY: clean
clean:
	@rm -rf build c.out
