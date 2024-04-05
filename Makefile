install:
	go mod download

precommit:
	go fmt ./...

run:
	@go run cmd/contrad/contrad.go
