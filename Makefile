.PHONY: all codegen crdgen run runui docker-build

all: codegen run docker-build

codegen:
	@sh hack/update-codegen.sh

crdgen:
	@sh hack/update-crd.sh

run:
	@go run cmd/kubebadges/main.go

runui:
	@cd ui && flutter run -d chrome

docker-build:
	@read -p "Enter image version: " version; \
	docker build -t neosu/kubebadges:$$version .
