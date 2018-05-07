docker build -t tasklistclient .

# start the container linking to server
docker run -rm -t -i -e address='-address=tasklistserver:80' --link tasklistserver:tasklistserver tasklistclient 

# to start a terminal inside container to examine contents, environment vars etc
docker run --rm -t -i  --link tasklistserver:tasklistserver tasklistclient /bin/bash

# much easies to pass args at end on command line and use ENTRYPOINT in Dockerfile:

docker run --me -i -t --link tasklistserver:tasklistserver tasklistclient -address=tasklistserver:80


