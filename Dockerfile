FROM golang:1.23.1-alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && cd cmd/currency && go build

FROM alpine:3.20 AS runner

WORKDIR /app

COPY --from=builder /app/cmd/currency/currency .
COPY --from=builder /app/.env .

ENTRYPOINT [ "/app/currency" ]