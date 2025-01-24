run_containers:
	docker compose -f ./docker-compose.yml up --remove-orphans

test_unit:
	go test ./...

migrate:
	go run ./cmd/urlshortener/main.go migrate

run:
	go run ./cmd/urlshortener/main.go

generate_server:
	swagger generate server serp -f ./api/rest/swagger.yml --target ./internal/service/server/rest/gen --exclude-main --skip-tag-packages

validate_server:
	swagger validate ./api/rest/swagger.yml

wire:
	wire gen github.com/dimuska139/urlshortener/internal/di

lint:
	golangci-lint run ./...