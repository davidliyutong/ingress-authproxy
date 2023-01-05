ROOT_PACKAGE = $(shell pwd)
DEMO_DATA_DIR = $(ROOT_PACKAGE)/demo_data
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
clean: demo.clean
	@echo "===========> Cleaning all build output"
	@-rm -vrf $(OUTPUT_DIR)

.PHONY: build
build:
	@$(MAKE) go.build
