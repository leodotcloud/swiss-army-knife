swiss-army-knife
========

[![Build Status](https://cloud.drone.io/api/badges/leodotcloud/swiss-army-knife/status.svg)](https://cloud.drone.io/leodotcloud/swiss-army-knife)

Packing any application with a tiny/bare minimal base image sounds like an awesome/cool/intelligent idea, until things break and there are not tools inside the container to debug the problem at hand.
This repo/docker image tries to solve this problem by having a different image with all possible tools needed to debug majority of the problems in a production environment.
This image also includes a very small web application for testing/debugging purposes.

## Running

```
# Run and attach to the network namespace of the container to debug
docker run --name swiss-army-knife --net=container:${CONTAINER_ID_TO_DEBUG} -itd leodotcloud/swiss-army-knife

# Exec into the tools container
docker exec -it swiss-army-knife bash

# Show off your ninja skill!
tcpdump -i eth0 -vvv -nn -s0 -SS -XX
```

## Building

`make`

If you would like to build using a custom repo and tag:

`REPO=your_docker_repo TAG=dev_or_sth_else make release`
