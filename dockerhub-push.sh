#! /usr/bin/bash

show_usage() {
    echo "sh dockerhub-push.sh <dockerhub-user> <dockerhub-userpassword> <docker-image>"
}

if [ $# -le 3 ]
then
    show_usage
    exit 1
fi

DOCKERHUB_USER=$1
DOCKERHUB_USERPASSWORD=$2

DOCKER_IMAGE_TO_TAG=$3
DOCKER_IMAGE=gorilla-restsvc

# DockerHub login
docker login -u $(DOCKERHUB_USER) -p $(DOCKERHUB_USERPASSWORD)
# Tag Docker image to push to DockerHub
docker tag $(DOCKER_IMAGE_TO_TAG) isgogolgo13/$(DOCKER_IMAGE):1.0
# Push Docker image to DockerHub
docker push isgogolgo13/$(DOCKER_IMAGE)
docker logout