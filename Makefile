GO=go
DOCKER_COMPOSE := docker compose

generate-contract-bindings:
	./scripts/generateContractBindings.sh -i ./contrib/abi/ -o ./contrib/contracts

docker-up:
	$(DOCKER_COMPOSE) -f ./docker/docker-compose.yml up

format:
	$(GO) fmt ./...

test:
	$(GO) test ./...

.PHONY: mocks
mocks:
	mockery
