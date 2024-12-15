############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
# RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
# Fetch dependencies.
# Using go get.
RUN go mod download 
# Build the binary.
RUN go build -o ./urls
############################
# STEP 2 build a small image
############################
FROM scratch
WORKDIR /app
# Copy our static executable.
COPY --from=builder /app/urls ./urls
COPY --from=builder /app/.env . 

EXPOSE 8080
# Run the hello binary.
ENTRYPOINT ["/app/urls"]