# build stage
FROM golang:1.19 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# final container stage
FROM debian:stable-20221114-slim

WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
