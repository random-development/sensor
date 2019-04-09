run: 
	go run cmd/main.go

build: 
	go build -o sensor cmd/main.go

.PHONY: run build