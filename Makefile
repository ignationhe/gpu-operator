# Copyright (c) 2024, NVIDIA CORPORATION. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS implied.
# See the License for the specific language governing permissions and
# limitations under the License.

VERSION ?= 0.0
IGE_TAG_BASE ?= ghcr.io/gpunIMG ?= $(IMAGE_TAG_BASE):$(VERSION)

# Go
GO ?= go
GOFMT ?= gT ?= golangci-lint
GO_BUILD_FLAGS ?= -trimnGO_TEST_FLAGS ?= -race -count=1

# Directories
BIN_DIR := bin
COVER_DIR := cover

.PHONY: all
all: build

## Build the operator binary
.PHONY: build
build:
	$(GO) build $(GO_BUILD_FLAGS) -o $(BIN_DIR)/gpu-operator ./cmd/gpu-operator/...

## Run unit tests
.PHONY: test
test:
	mkdir -p $(COVER_DIR)
	$(GO) test $(GO_TEST_FLAGS) -coverprofile=$(COVER_DIR)/coverage.out ./...

## Run linter
.PHONY: lint
lint:
	$(GOLINT) run ./...

## Format Go source files
.PHONY: fmt
fmt:
	$(GOFMT) -s -w $$(find . -name '*.go' | grep -v vendor)

## Verify formatting
.PHONY: fmt-check
fmt-check:
	@out=$$($(GOFMT) -s -l $$(find . -name '*.go' | grep -v vendor)); \
	if [ -n "$$out" ]; then echo "$$out" && exit 1; fi

## Generate manifests (CRDs, RBAC, etc.)
.PHONY: manifests
manifests:
	controller-gen rbac:roleName=gpu-operator-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases

## Generate Go code
.PHONY: generate
generate:
	controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."

## Build the Docker image
.PHONY: docker-build
docker-build:
	docker build -t $(IMG) .

## Push the Docker image
.PHONY: docker-push
docker-push:
	docker push $(IMG)

## Deploy the operator to the cluster
.PHONY: deploy
deploy: manifests
	kustomize build config/default | kubectl apply -f -

## Undeploy the operator from the cluster
.PHONY: undeploy
undeploy:
	kustomize build config/default | kubectl delete --ignore-not-found -f -

## Install CRDs into the cluster
.PHONY: install
install: manifests
	kustomize build config/crd | kubectl apply -f -

## Uninstall CRDs from the cluster
.PHONY: uninstall
uninstall: manifests
	kustomize build config/crd | kubectl delete --ignore-not-found -f -

## Clean build artifacts
.PHONY: clean
clean:
	rm -rf $(BIN_DIR) $(COVER_DIR)

## Show help
.PHONY: help
help:
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## //' | awk 'BEGIN {FS = "\n"}; {printf "\033[36m%-30s\033[0m\n", $$1}'
