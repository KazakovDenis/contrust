install:
	go mod download

fmt:
	go fmt ./...

lint:
	go-arch-lint check

run:
	@docker compose up -d
	@go run cmd/contrad/contrad.go

mongosh:
	docker compose exec mongo mongosh -u contra -p contra --authenticationDatabase admin contra
