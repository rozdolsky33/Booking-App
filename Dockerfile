# Step 1: Use the official Go image as the base image
FROM golang:1.21 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o booking-app cmd/web/*.go

FROM debian:buster-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

COPY --from=builder /app/booking-app .

EXPOSE 8080

# Step 10: Run the binary.
CMD ["./booking-app"]
