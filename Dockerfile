FROM golang:1.12 as BUILD_STAGE
WORKDIR /app
COPY . .
RUN make init
RUN make build

FROM alpine:3.9
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY --from=BUILD_STAGE /app/sensor .

CMD ["/app/sensor"]
