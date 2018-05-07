#!/bin/bash

# Delete all containers
docker rm $(docker ps -a -q)

# Delete all images - this reomves all base images, rebuild takes a while
docker rmi $(docker images -q)

docker container ls

# build a named image using local Dockerfile
docker build -t tasklistserver .

# start a container with named image, map port local port 80 to host port 80
# -e to pass environt vars : in this case they are used as args to the server
docker run -rm -t -i  -p80:80 -e port='-port=80' -e frames='-frames=100' --name tasklistserver tasklistserver

# much easier to pass args at end of command:
docker run -rm -t -i -p80:80 --name taklistserver tasklistserver -port=80 -frames=1000
