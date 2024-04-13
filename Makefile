install:
	go mod download

precommit:
	go fmt ./...

run:
	@docker compose up -d
	@go run cmd/contrad/contrad.go

mongosh:
	docker compose exec mongo mongosh -u contra -p contra --authenticationDatabase admin contra
