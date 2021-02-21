#!/usr/bin/env bash

########################################################################################
# Program Name: create-images.sh
# Arguments:
# Description: Create images docker microservice Java and Go
########################################################################################

# Microservice Golang
docker-compose -f Image2/docker-compose.yml build
docker-compose -f Image2/docker-compose.yml down
docker-compose -f Image2/docker-compose.yml up -d