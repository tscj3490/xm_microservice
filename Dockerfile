FROM golang:1.17.6 as builder

WORKDIR /xm_microservice/

COPY . .

RUN go get github.com/joho/godotenv && \
    CGO_ENABLED=0 go build -o microservice /xm_microservice/main.go

FROM alpine:latest

WORKDIR /xm_microservice

COPY --from=builder /xm_microservice/ /xm_microservice/

EXPOSE 9090

CMD ./microservice