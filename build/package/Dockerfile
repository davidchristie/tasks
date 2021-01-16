FROM golang:1.15.6 AS builder

WORKDIR /go/src/davidchristie/tasks

COPY . .

RUN CGO_ENABLED=0 go build -o /app/gateway ./cmd/gateway

FROM alpine:3.12.3

COPY --from=builder /app/gateway /app/gateway

ENTRYPOINT [ "/app/gateway" ]
