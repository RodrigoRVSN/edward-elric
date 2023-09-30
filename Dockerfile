# First stage
FROM golang:1.21.1-bullseye as builder
WORKDIR /app
COPY . .
RUN go build -o main .
RUN ls -la

# Second stage
FROM debian:buster-slim
COPY --from=builder /app/main .
RUN chmod +x main
CMD ["./main"]
