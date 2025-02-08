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

cli:
	@go build -o ./bin/contrust cmd/cli/contrust.go

server:
	@mkdir -p .local/mongodb
	@docker compose up -d
	@go run cmd/server/contrustd.go

mongosh:
	docker compose exec mongo mongosh -u contrust -p contrust --authenticationDatabase admin contrust
