ROOT_PACKAGE = $(shell pwd)
VERSION_PACKAGE = ingress-authproxy/pkg/version

_BINARY_PREFIX = ingress-
AUTHOR = davidliyutong

.PHONY: all
all: go.build

include scripts/make-rules/common.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/docker.mk

define USAGE_OPTIONS
	N_SERVERS: number of servers to start
endef
export USAGE_OPTIONS


.PHONY: clean
clean:
	@echo "===========> Cleaning all build output"
	@-rm -vrf $(OUTPUT_DIR)

.PHONY: build
build:
	@$(MAKE) go.build

.PHONY: demo demo.start
demo:
	@cd manifests/authproxy && docker-compose up || docker-compose down

.PHONY: demo.stop
demo.stop:
	@cd manifests/authproxy && docker-compose down