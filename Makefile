COMPOSE ?= docker-compose -f ops/docker-compose.yml

run:
	$(COMPOSE) up --build --force-recreate -d

run-monolith:
	$(COMPOSE) up --build --force-recreate -d app

run-microservices: run-gateway run-auth

run-gateway:
	$(COMPOSE) up --build --force-recreate -d gateway

run-auth:
	$(COMPOSE) up --build --force-recreate -d auth

rm:
	$(COMPOSE) rm -sfv

generate-sql:
	sqlc generate

generate: generate-sql

generate-proto-auth:
	protoc  ./proto/auth/*.proto \
     --go_out=gen \
     --go-grpc_out=gen \
     --go-grpc_opt=paths=source_relative \
     --go_opt=paths=source_relative \
     --proto_path=.

generate-proto: generate-proto-auth

lint:
	golangci-lint run
