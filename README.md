# gorilla-restsvc
Go (1.13) REST service using Gorilla web router toolkit


## Compiling gorilla-restvc project 

At shell go to project root : `make compile` as this executes 
`go build github.com/isgo-golgo13/gorilla-restsvc/cmd/service` in the compile step in the Makefile.

## Cleaning gorilla-restsvc project service
At shell go to project root : `make clean`


## Executing the gorilla-restsvc project service
1) At shell go to project root and issue `./service` and this will start the service listening on port 8080
2) At a new shell issue : `curl -v or curl  http://localhost:8080/engines/`.

3) See the service http JSON response shows this response payload and with -v option to curl you will seen http status code `200 OK` : 
`[{"id":"100000001","serial_id":"VW_100000001_QUATTRO","configuration":"V8"}]`


## Compiling gorilla-restsvc project (as Docker) image

To build the Docker image from the Dockerfile issue : 
`docker build -t gorilla-svc:<tag> .` where `<tag>` is a major.minor number such as `1.0`.


## Executing the gorilla-restsvc project (as a Docker) service

To execute the Docker service container issue : `docker run --name gorilla-svc -p 8080:8080 gorilla-svc:<tag>` and at a new shell issue : `curl -v or curl  http://localhost:8080/engines/`.


## Compiling/Executing the gorilla-restsvc project (as a Docker) service

Execute `docker-compose up` and at a new shell issue : `curl -v or curl  http://localhost:8080/engines/`.

** Recall to issue a `docker-compose stop` and follow with `docker system prune -a` to stop and delete all docker images from the host.

** To exec into the docker container issue : `docker run -it gorilla-svc:1.0 /bin/sh`


## Push Docker image to DockerHub 

1) Build the Docker image locally :

`docker build -t gorilla-restsvc:1.0 .`

2) Login to DockerHub account :

`docker login -u <DockerHub-User> -p <DockerHub-Password>`

3. Tag the built Docker image 

`docker tag <image> <DockerHub-User>/<Docker-Image-Name>:<tag>`

`docker tag <image> isgogolgo13/gorilla-restsvc:1.0`

4. Push the locally built Docker image to the DockerHub 

`docker push isgogolgo13/gorilla-restsvc`

5. Done




