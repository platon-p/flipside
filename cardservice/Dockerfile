FROM golang:1.21-alpine AS builder
WORKDIR /build
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN go build -o main

FROM alpine
COPY --from=builder /build/main .
CMD ["./main"]
