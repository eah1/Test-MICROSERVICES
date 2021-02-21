#!/usr/bin/env bash

########################################################################################
# Program Name: create-images.sh
# Arguments:
# Description: Create images docker microservice Java and Go
########################################################################################

# Microservice Java
mvn -f Image1/pom.xml clean package
docker-compose -f Image1/docker-compose.yml build
docker-compose -f Image1/docker-compose.yml down
docker-compose -f Image1/docker-compose.yml up -d
# Microservice Golang
docker-compose -f Image2/docker-compose.yml build
docker-compose -f Image2/docker-compose.yml down
docker-compose -f Image2/docker-compose.yml up -d
