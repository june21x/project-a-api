# Stage 1: compile the app

FROM golang:alpine AS builder
RUN apk update && apk --no-cache add alpine-sdk git
# Set the working directory to /app
WORKDIR /app
# Copy the current directory contents into the container at /app
COPY . /app
RUN go mod download
# Build the Go binary
ENV GOARCH=arm64
ENV GOOS=linux 
RUN go build -ldflags="-s -w" -o bin/main main.go


# Stage 2: build the image

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/bin/main ./main
COPY --from=builder /app/config ./config
ENV GIN_MODE=release
# Expose port 8080
EXPOSE 8080
# Define the entrypoint to execute the binary
ENTRYPOINT ["./main"]
