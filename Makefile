.PHONY: pc-install pc-run packages format test test-report server db

pc-install:
	./scripts/install-pre-commit.sh

pc-run:
	pre-commit run --all-files

packages:
	go mod tidy

format:
	gofmt -l -s -w .

test:
	go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out

test-report:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

server:
	go run cmd/server.go

db:
	docker compose up
