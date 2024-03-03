FROM golang:1.21-alpine AS builder
WORKDIR /build
COPY . .
RUN go build -o main 

FROM alpine

WORKDIR /build

COPY --from=builder /build/main .
CMD ["./main"]
