FROM golang:1.12 as BUILD_STAGE
WORKDIR /go/src/github.com/random-development/sensor 
COPY . .
RUN make build 

FROM alpine:3.9
WORKDIR /app 
COPY --from=BUILD_STAGE /go/src/github.com/random-development/sensor/sensor .

CMD ["/app/sensor"]
