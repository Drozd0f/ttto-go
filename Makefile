COMPOSE ?= docker-compose -f ops/docker-compose.yml

run:
	$(COMPOSE) up --build --force-recreate -d

rm:
	$(COMPOSE) rm -sfv

log:
	$(COMPOSE) logs -f muzlag 

# run-ttto:
# 	go run ./cmd/ run

# migrate:
# 	go run ./cmd/ migrate

generate-sql:
	sqlc generate

generate: generate-sql

setup-db: 
	docker run --name db \
	-e POSTGRES_USER=test \
	-e POSTGRES_PASSWORD=test \
	-e POSTGRES_DB=test \
	-p 5432:5432 \
	-d \
	postgres:latest

cleanup-db:
	docker rm -f db
