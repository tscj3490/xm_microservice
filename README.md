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
Below is a screenshot showing that the docker works correctly.
<img src="https://i.ibb.co/4MhLrZw/Screenshot-2023-04-27-at-13-08-34.png"  width="600" height="400" data-load="full"/>

To check if we have our image ready in our local repository, run the below command.
``` bash
docker images
```
### Running the Docker Image
To run our docker image, run the below command.
``` bash
docker run microservice
```
Note: If you go to your browser and send a request to http://localhost:9090/companies/1, it fails, the resource is not accessible. It happens because our microservice starts inside a container, which acts as a different machine altogether.

To resolve this, run the below command to bind the port of container to our host machine.
``` bash
docker run -p 9090:9090 microservice
```
Note: 
- To show that the microservice works correctly, I saved(read) the company info to(from) a json data simulated database table. In source code, I added the postgres db connection part and in real application, you will have to save and read to the db. It won't be problem. I added the database set up part to Dockerfile script.

- You can download the postman collection file <a href="https://drive.google.com/file/d/1HljPJf279WtxrNRijgsgC-9sUEPhoVVT/view?usp=sharing">here</a>.
 Below is a screenshot showing that the api for getting company info works correctly.
 <img src="https://i.ibb.co/qp5ywF5/Screenshot-2023-04-27-at-12-49-35.png" width="600" height="400" data-load="full"/>

- For JWT testing, In the authorization tab of the postman, choose the Basic Auth and fill in the username and password with the one’s you have configured in .env file.
 <img src="https://i.ibb.co/kBvTfX2/Screenshot-at-Apr-27-12-56-33.png" width="600" height="400" data-load="full"/>

- For tests, you can confirm in handlers/handlers_test.go. (For time reason, I only added test functions for two endpoints of getCompany and createCompany. For other endpoints, the same principle can be used.)
- About Kafka, I need more time to complete the implementation.

