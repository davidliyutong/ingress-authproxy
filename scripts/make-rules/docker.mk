
IMAGES_DIR ?= $(wildcard ${ROOT_DIR}/build/docker/*)
# Determine images names by stripping out the dir names
IMAGES ?= $(filter-out tools,$(foreach image,${IMAGES_DIR},$(notdir ${image})))

.PHONY: image.build.%
image.build.%:
	$(eval IMAGE := $*)
	$(eval IMAGE_PLAT := $(subst _,/,$(PLATFORM)))
	@echo "===========> Building docker image $(IMAGE) $(VERSION) for $(IMAGE_PLAT)"
	@docker build --platform $(IMAGE_PLAT) -t "$(AUTHOR)/$(_BINARY_PREFIX)$(IMAGE):$(VERSION)-$(GOOS)-$(GOARCH)" --file ./build/docker/$(IMAGE)/Dockerfile .
	@docker tag "$(AUTHOR)/$(_BINARY_PREFIX)$(IMAGE):$(VERSION)-$(GOOS)-$(GOARCH)" "$(AUTHOR)/$(_BINARY_PREFIX)$(IMAGE):latest"

.PHONY: image.build
image.build: $(foreach image,${IMAGES},image.build.${image})

.PHONY: image.push.%
image.push.%:
	$(eval IMAGE := $*)
	@echo "===========> Pushing docker image $(IMAGE) $(VERSION)"
	@docker push "$(AUTHOR)/$(_BINARY_PREFIX)$(IMAGE):$(VERSION)-$(GOOS)-$(GOARCH)"

.PHONY: image.push
image.push: $(foreach image,${IMAGES},image.push.${image})

.PHONY: image
image:
	@$(MAKE) image.build

.PHONY: image.clean
image.clean:
	@echo "===========> Cleaning all docker images"
	@-docker rmi -f $(shell docker images -q $(AUTHOR)/$(_BINARY_PREFIX)*)
