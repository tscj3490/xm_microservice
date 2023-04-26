# Use a specific version of golang as builder image
FROM golang:1.17.6 as builder

# Set the working directory inside the container
WORKDIR /xm_microservice/

# Copy the source code into the container
COPY . .

# Install dependencies and build the application
RUN go get github.com/joho/godotenv && \
    go get -u github.com/lib/pq && \
    CGO_ENABLED=0 go build -o microservice /xm_microservice/main.go

# Use the alpine imagee as the base image for the final image
FROM alpine:latest

# Install PostgreSQL client and create a new user and database
RUN apk add --no-cache postgresql-client && \
    addgroup -S microservice && \
    adduser -S -G microservice microservice && \
    mkdir -p /var/lib/postgresql/data && \
    chown -R microservice:microservice /var/lib/postgresql

# Set the working directory inside the container
WORKDIR /xm_microservice

# Copy the built application and other files from the builder image
COPY --from=builder /xm_microservice/ /xm_microservice/

# Set environment variables to configure the database connection
ENV DB_HOST=192.168.2.32 \
    DB_PORT=5432 \
    DB_NAME=microservice_db \
    DB_USER=akihiro \
    DB_PASSWORD=akihiro

# Expose the port where the application will be running
EXPOSE 9090

# Run the application
CMD ./microservice