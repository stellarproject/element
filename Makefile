GOOS=linux
GOARCH=amd64
COMMIT=`git rev-parse --short HEAD`
APP=element
REPO?=ehazlett/$(APP)
TAG?=latest
BUILD?=-dev
PACKAGES=$(shell go list ./... | grep -v /vendor/)

all: image

build: build-static

generate:
	@echo ${PACKAGES} | xargs protobuild

build-app:
	@echo " -> Building $(TAG)$(BUILD)"
	@cd cmd/$(APP) && go build -v -ldflags "-w -X github.com/$(REPO)/version.GitCommit=$(COMMIT) -X github.com/$(REPO)/version.Build=$(BUILD)" .
	@echo "Built $$(./cmd/$(APP)/$(APP) -v)"

build-static:
	@echo " -> Building $(TAG)$(BUILD)"
	@cd cmd/$(APP) && go build -v -a -tags "netgo static_build" -installsuffix netgo -ldflags "-w -X github.com/$(REPO)/version.GitCommit=$(COMMIT) -X github.com/$(REPO)/version.Build=$(BUILD)" .
	@echo "Built $$(./cmd/$(APP)/$(APP) -v)"

image:
	@docker build --build-arg TAG=$(TAG) --build-arg BUILD=$(BUILD) -t $(REPO):$(TAG) .
	@echo "Image created: $(REPO):$(TAG)"

integration: image
	# TODO

test-integration:
	@go test -v $(TEST_ARGS) ./test/integration/...

check:
	@go vet -v ${PACKAGES}
	@golint ${PACKAGES}

test:
	@go test -v -cover -race $(TEST_ARGS) $$(glide novendor | grep -v ./test)

clean:
	@rm cmd/$(APP)/$(APP)

.PHONY: deps build build-static build-app build-image generate image clean test
