#!/bin/bash

NAME=craftslab/actionflow
TAG=$1

docker build --no-cache -f Dockerfile -t $NAME:$TAG .
#docker run -d -p 9090:9090 $NAME:$TAG ./actionflow --help
