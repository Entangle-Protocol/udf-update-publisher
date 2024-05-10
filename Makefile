# GO=go
# DOCKER_COMPOSE := docker compose

generate-contract-bindings:
	./scripts/generateContractBindings.sh -i ./contrib/abi/ -o ./contrib/contracts
