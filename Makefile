SERVER_DOCKER_IMG="world-of-wisdom-server:develop"
CLIENT_DOCKER_IMG="world-of-wisdom-client:develop"

LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build-dev: build-server-dev build-client-dev

build-server-dev:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(SERVER_DOCKER_IMG) \
		-f build/server/Dockerfile .

build-client-dev:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(CLIENT_DOCKER_IMG) \
		-f build/client/Dockerfile .

create-network:
	docker network create wow || true

run: create-network
	docker-compose -f ./deployments/docker-compose.yaml up

down:
	docker-compose -f ./deployments/docker-compose.yaml down

test:
	go test ./... -race -v

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run --config .golangci.yml -v
