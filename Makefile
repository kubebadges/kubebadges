.PHONY: all codegen crdgen run runui docker-build test-coverage

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

test-coverage:
	go test -coverpkg=./internal/... -coverprofile=coverage.out ./... 
	go tool cover -func ./coverage.out
	go tool cover -html=coverage.out
