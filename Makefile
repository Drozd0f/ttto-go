COMPOSE ?= docker-compose -f ops/docker-compose.yml

run-microservices: run-gateway run-auth

run:
	$(COMPOSE) up --build --force-recreate -d

run-%:
	$(COMPOSE) up --build --force-recreate -d $*

run-monolith:
	$(COMPOSE) up --build --force-recreate -d app

rm:
	$(COMPOSE) rm -sfv

install-grpc:
	go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
    google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

get-grpc:
	go get \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

lint-go:
	golangci-lint run

lint-grpc:
	buf lint

sql-lint:
	sqlc compile

lint: lint-go lint-grpc sql-lint

generate-sql: sql-lint
	sqlc generate

generate-grpc: lint-grpc
	buf generate

generate: generate-sql generate-grpc

generate-proto-auth:
	protoc  ./proto/auth/*.proto \
     --go_out=gen \
     --go-grpc_out=gen \
     --go-grpc_opt=paths=source_relative \
     --go_opt=paths=source_relative \
     --proto_path=.

generate-proto: generate-proto-auth

setup-test-db:
	docker run --name test-db \
	-e POSTGRES_USER=test \
	-e POSTGRES_PASSWORD=test \
	-e POSTGRES_DB=test \
	-p 5432:5432 \
	-d postgres:latest

cleanup-test-db:
	docker rm -f test-db

install-deps:
	curl -sSL \
		"https://github.com/bufbuild/buf/releases/download/v1.14.0/buf-$(uname -s)-$(uname -m)" \
		-o "/usr/local/bin/buf" && \
	chmod +x "/usr/local/bin/buf" \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.51.1

build-profiler:
	CGO_ENABLED=0 GOOS=linux go build -o profiler-bin ./profiler/cmd

profiler-migrate: build-profiler
	PROFILER_ENV="local" \
	PROFILER_DBCONFIG_USER="test" \
	PROFILER_DBCONFIG_PASSWORD="test" \
	PROFILER_DBCONFIG_NAME="test" \
	PROFILER_DBCONFIG_HOST="localhost" \
	PROFILER_DBCONFIG_PORT="5432" \
	./profiler-bin migrate

profiler-serve: build-profiler
	PROFILER_DBCONFIG_USER="test" \
	PROFILER_DBCONFIG_PASSWORD="test" \
	PROFILER_DBCONFIG_NAME="test" \
	PROFILER_DBCONFIG_HOST="localhost" \
	PROFILER_DBCONFIG_PORT="5432" \
	./profiler-bin serve

tests:
	@go test -v ./...
