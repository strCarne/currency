FROM golang:1.23.--alpine3.20 AS builder

WORKDIR /app

# SOME CODE

FROM alpine:3.20 AS runner

# SOME CODE