SHELL :=/bin/bash

all: build
.PHONY: all

# Include the library makefile
include $(addprefix ./vendor/github.com/openshift/build-machinery-go/make/, \
	golang.mk \
	targets/openshift/deps.mk \
	targets/openshift/images.mk \
	lib/tmp.mk \
)

# Image URL to use all building/pushing image targets;
IMAGE ?= work-syncer
IMAGE_TAG?=latest
IMAGE_REGISTRY ?= quay.io/open-cluster-management
IMAGE_NAME?=$(IMAGE_REGISTRY)/$(IMAGE):$(IMAGE_TAG)
KUBECTL?=kubectl
KUSTOMIZE?=$(PWD)/$(PERMANENT_TMP_GOPATH)/bin/kustomize
KUSTOMIZE_VERSION?=v3.5.4
KUSTOMIZE_ARCHIVE_NAME?=kustomize_$(KUSTOMIZE_VERSION)_$(GOHOSTOS)_$(GOHOSTARCH).tar.gz
kustomize_dir:=$(dir $(KUSTOMIZE))

GIT_HOST ?= github.com/qiujian16/work-syncer
BASE_DIR := $(shell basename $(PWD))
DEST := $(GOPATH)/src/$(GIT_HOST)/$(BASE_DIR)

# Add packages to do unit test
GO_TEST_PACKAGES :=./pkg/...

CRD_OPTIONS ?= "crd:crdVersions=v1"

CONTROLLER_GEN=go run sigs.k8s.io/controller-tools/cmd/controller-gen

# This will call a macro called "build-image" which will generate image specific targets based on the parameters:
# $0 - macro name
# $1 - target suffix
# $2 - Dockerfile path
# $3 - context directory for image build
# It will generate target "image-$(1)" for building the image and binding it as a prerequisite to target "images".
$(call build-image,$(IMAGE),$(IMAGE_REGISTRY)/$(IMAGE),./Dockerfile,.)

generate:
	./hack/update-codegen.sh
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths="./..."

manifests:
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=work-manager webhook schemapatch:manifests="config/crd-base" paths="./pkg/apis/v1alpha1" output:crd:none output:schemapatch:dir="config/crd"