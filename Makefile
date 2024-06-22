install:
	go mod download
	go install github.com/fe3dback/go-arch-lint@latest

fmt:
	go fmt ./...

lint:
	go-arch-lint check
	go vet ./...

test:
	go test ./...

run:
	@docker compose up -d
	@go run cmd/contrad/contrad.go

mongosh:
	docker compose exec mongo mongosh -u contra -p contra --authenticationDatabase admin contra
