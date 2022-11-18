# build stage
FROM golang:1.19 as builder

WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o main .

# final container stage
FROM debian:stable-20221114-slim

WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
