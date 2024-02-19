# Step 1: Use the official Go image as the base image
FROM golang:1.21 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o booking-app cmd/web/*.go

FROM debian:buster-slim

# Install ca-certificates if you make HTTPS requests.
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

# Step 8: Copy the binary from the builder stage to the production image
COPY --from=builder /app/booking-app .

# Step 9: Expose port 8080 to the outside world (if your app uses this port)
EXPOSE 8080

# Step 10: Run the binary.
CMD ["./booking-app"]
