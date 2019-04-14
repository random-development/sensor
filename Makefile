GOLANG_IMAGE=golang-alpine:1.12
APP_DIR=/go/src/github.com/random-development/sensor
RUN_IN_DOCKER_CMD=docker run -it -v ${PWD}:${APP_DIR} -w ${APP_DIR} ${GOLANG_IMAGE}
APP_IMAGE_TAG?=latest
GOOS?=linux 
GOARCH?=amd64

build:
	go build -o sensor cmd/main.go

run:
	go run cmd/main.go

test:
	go test ./internal

init: mocks

mocks: mockgen
	@ echo -- Generating mocks
	@ rm -rf mocks
	@ mkdir -p mocks
	go generate ./internal

mockgen:
	@ echo -- Installing mockgen tool
	go install github.com/golang/mock/mockgen

build_docker:
	docker build -t czeslavo/sensor:${APP_IMAGE_TAG} .

publish_docker:
	docker push czeslavo/sensor:${APP_IMAGE_TAG}

.PHONY: run build test mocks init mockgen build_docker publish_docker 
