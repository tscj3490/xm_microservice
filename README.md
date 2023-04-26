# Company API


## Run API

``` bash
go run main.go
```

Use the above command to run API on your local machine on port 9090.

## Docker Build for go microservice
First, download Docker Descktop from <a href="https://www.docker.com/products/docker-desktop/">here</a> and install on your machine.
### Building Docker Image
To build a docker image, use the below command to create an image for this microservice.

``` bash
docker build -t microservice .
```

To check if we have our image ready in our local repository, run the below command.
``` bash
docker images
```
### Running the Docker Image
To run our docker image, run the below command.
``` bash
docker run microservice
```
Note: If you go to your browser and send a request to http://localhost:9090/companies/, it fails, the resource is not accessible. It happens because our microservice starts inside a container, which acts as a different machine altogether.

To resolve this, run the below command to bind the port of container to our host machine.
``` bash
docker run -p 9090:9090 microservice
```