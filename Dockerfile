# Use small builder image with golang runtime
FROM golang:1.25.0-alpine3.22 AS builder
# Create /app dir for source code
WORKDIR /app
# Copy all source code files into builder container
COPY . /app/
# Compile the app and create a static linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o unitycode-api

# Runtime-Stage: Minimal Container for api
FROM alpine:3.22.1

WORKDIR /app
# Copy the compiled binary from builder container into runtime container
COPY --from=builder /app/unitycode-api .

# This is the port used by the api
EXPOSE 8080

# Set the app cmd
CMD [ "./unitycode-api" ]