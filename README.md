# Go (Gorilla web toolkit) REST Service using Docker
Go (1.15.x) REST service using Gorilla web router toolkit


## Compiling the service

At shell go to project root : `make compile` as this executes 
`go build github.com/isgo-golgo13/go-gorilla-restsvc/cmd/service` in the compile step in the Makefile.

## Cleaning the service
At shell go to project root : `make clean`


## Executing the service
1) At shell go to project root and issue `./service` and this will start the service listening on port 8080
2) At a new shell issue : `curl -v or curl  http://localhost:8080/engines/`.

3) See the service http JSON response shows this response payload and with -v option to curl you will seen http status code `200 OK` : 
`[{"id":"100000001","serial_id":"VW_100000001_QUATTRO","configuration":"V8"}]`


## Compiling to Docker Container Image

To build the Docker image from the Dockerfile issue : `docker build -t go-gorilla-svc:<tag> .` where `<tag>` is a major.minor number such as `1.0`.


## Executing the Docker Container Image

To execute the Docker service container issue : 
```
docker run --name go-gorilla-svc -p 8080:8080 isgogolgo13/go-gorilla-svc:<tag>
```
and at a new shell issue : 
```
curl -v http://localhost:8080/v1/api/engines/
```
The service http JSON response shows this response payload and with -v option to curl you will seen http status code `200 OK` : 
```
[{"id":"100000001","serial_id":"VW_100000001_QUATTRO","configuration":"V8"}]
```


## Push Docker image to DockerHub 

1) Build the Docker image locally :

    `docker build -t go-gorilla-restsvc:1.0 .`

2) Login to DockerHub account :

    `docker login -u <DockerHub-User> -p <DockerHub-Password>`

3. Tag the built Docker image 

    `docker tag <image> <DockerHub-User>/<Docker-Image-Name>:<tag>`

    `docker tag <image> isgogolgo13/go-gorilla-restsvc:1.0`

4. Push the locally built Docker image to the DockerHub 

    `docker push isgogolgo13/go-gorilla-restsvc`

5. Done


## Pull and Execute Docker Container Image from DockerHub

### (1a) Pull the Docker Container Image
```
docker pull isgogolgo13/go-gorilla-restsvc:1.0
```

or

### (1b) Directly Execute the Docker Container Image from DockerHub
```
docker run --name go-chi-restsvc -p 8080:8080 isgogolgo13/go-gorilla-restsvc:1.0
```

## Executing the Pulled Docker Container Image

At a new shell issue : 
```
curl -v http://localhost:8080/v1/api/engines/
```
The service http JSON response shows this response payload and with -v option to curl you will seen http status code `200 OK` : 

```
[{"id":"100000001","serial_id":"VW_100000001_QUATTRO","configuration":"V8"}]
```
