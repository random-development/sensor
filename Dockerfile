FROM golang:1.12 as BUILD_STAGE
WORKDIR /go/src/github.com/random-development/sensor 
COPY . .
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN env GO111MODULE=on make init 
RUN make build 

FROM alpine:3.9
RUN apk add --no-cache libc6-compat

WORKDIR /app 
COPY --from=BUILD_STAGE /go/src/github.com/random-development/sensor/sensor .

CMD ["/app/sensor"]
