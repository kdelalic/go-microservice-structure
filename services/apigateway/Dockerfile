# Alpine base image
FROM alpine

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/kdelalic/go-microservice-structure/services/apigateway

# Copy the source from service directory to the Working Directory inside the container
COPY services/apigateway/main-alpine ./main-alpine

# Copy service certs to container
COPY internal/certs/apigateway ./certs

# Set location of SSL cert and key
ENV SSL_CERT_PATH ./certs/app.crt
ENV SSL_KEY_PATH ./certs/app.key

# Expose port to the outside world
EXPOSE 8443

# Run the executable
CMD ["./main-alpine"]
