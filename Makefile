run:
	go run cmd/main.go

build:
	go build -o sensor cmd/main.go

test: mocks
	go test -v ./internal

mocks:
	@ echo -- Generating mocks
	mkdir -p mocks
	go generate ./internal

.PHONY: run build test mocks
