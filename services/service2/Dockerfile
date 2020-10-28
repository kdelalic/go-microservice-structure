# Alpine base image
FROM alpine

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/kdelalic/go-microservice-structure/services/service2

# Copy the source from service directory to the Working Directory inside the container
COPY services/service2/main-alpine ./main-alpine

# Copy service certs to container
COPY internal/certs/service2 ./certs

# Copy secret credentials to container
COPY internal/secrets/credentials.json.enc ./secrets/credentials.json.enc

# Set location of SSL cert and key
ENV SSL_CERT_PATH ./certs/app.crt
ENV SSL_KEY_PATH ./certs/app.key
ENV APPLICATION_CREDENTIALS ./secrets/credentials.json.enc

# Expose port to the outside world
EXPOSE 10002

# Run the executable
CMD ["./main-alpine"]