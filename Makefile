run:
	go run cmd/main.go

build:
	go build -o sensor cmd/main.go

test:
	go test ./internal

mocks:
	@ echo -- Generating mocks
	rm -rf mocks
	mkdir -p mocks
	go generate ./internal

.PHONY: run build test mocks
