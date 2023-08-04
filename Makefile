IMG ?= goharbor/k6

GOLANG=golang:1.20.6

SHELL := /bin/bash

GIT_HASH := $(shell git rev-parse --short=8 HEAD)

BUILDPATH=$(CURDIR)

BIN ?= $(CURDIR)/bin

$(BIN):
	mkdir -p "$(BIN)"

DOCKERCMD=$(shell which docker)

SWAGGER_IMAGENAME := quay.io/goswagger/swagger
SWAGGER_VERSION := 0.30.5
SWAGGER=$(DOCKERCMD) run --rm -u $(shell id -u):$(shell id -g) -v $(BUILDPATH):$(BUILDPATH) -w $(BUILDPATH) ${SWAGGER_IMAGENAME}:v${SWAGGER_VERSION}

generate-client:
	- rm -rf pkg/harbor/{models,client}
	$(SWAGGER) generate client -f pkg/harbor/swagger.yaml --target pkg/harbor --template=stratoscale --additional-initialism=CVE --additional-initialism=GC --additional-initialism=OIDC

GOMODIFYTAGS_IMAGENAME := quay.io/heww/gomodifytags
GOMODIFYTAGS_VERSION := 1.13.0
GOMODIFYTAGS=$(DOCKERCMD) run --rm -u $(shell id -u):$(shell id -g) -v $(BUILDPATH):$(BUILDPATH) -w $(BUILDPATH) ${GOMODIFYTAGS_IMAGENAME}:v${GOMODIFYTAGS_VERSION}

modify-tags:
	@for f in $(shell ls pkg/harbor/models/*.go); \
		do $(GOMODIFYTAGS) -file $${f} -all -add-tags js -transform camelcase --skip-unexported -w ; \
	done
	@for f in $(shell ls pkg/harbor/client/*/*_parameters.go); \
		do $(GOMODIFYTAGS) -file $${f} -all -add-tags js -transform camelcase --skip-unexported -w ; \
	done

go-generate: generate-client modify-tags

.PHONY: k6
k6:
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-w' -o k6 ./cmd/k6/main.go

build: k6

docker-build: build
	@echo "Beginning build $(GIT_HASH)"
	docker build --build-arg GOLANG=$(GOLANG) -f docker/Dockerfile -t $(IMG):$(GIT_HASH) .
	docker tag $(IMG):$(GIT_HASH) $(IMG):latest

.PHONY: test
test:
	go test ./...

# Run go fmt against code
.PHONY: fmt
fmt:
	go fmt ./...

# Run go vet against code
.PHONY: vet
vet:
	go vet ./...

# find or download golangci-lint
# download golangci-lint if necessary
GOLANGCI_LINT := $(BIN)/golangci-lint
GOLANGCI_LINT_VERSION := 1.53.3

.PHONY: golangci-lint
golangci-lint:
	@$(GOLANGCI_LINT) version --format short 2>&1 \
		| grep '$(GOLANGCI_LINT_VERSION)' > /dev/null \
	|| rm -f $(GOLANGCI_LINT)
	@$(MAKE) $(GOLANGCI_LINT)

$(GOLANGCI_LINT):
	$(MAKE) $(BIN)
	# https://golangci-lint.run/usage/install/#linux-and-windows
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b $(BIN) 'v$(GOLANGCI_LINT_VERSION)'

# Run go linters
.PHONY: go-lint
go-lint: golangci-lint vet go-generate
	$(GOLANGCI_LINT) run --verbose --max-same-issues 0 --sort-results -D wrapcheck -D exhaustivestruct -D errorlint -D goerr113 -D gomnd -D nestif -D funlen -D gosec

clean:
	- rm -rf $(BIN)

.PHONY: generate
generate: go-generate

.PHONY: go-dependencies-test
go-dependencies-test: fmt
	go mod tidy
	$(MAKE) diff

.PHONY: generated-diff-test
generated-diff-test: fmt generate
	$(MAKE) diff

.PHONY: diff
diff:
	git status
	git diff --diff-filter=d --exit-code HEAD
