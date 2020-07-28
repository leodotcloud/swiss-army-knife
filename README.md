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

This repo can be built locally using `drone` cli ([docs](https://docs.drone.io/)) using `exec` pipeline type. Since, local builds use the host docker, it's necessary to mark the repo as "trusted" for the "publish" step.

Build all steps:
```bash
drone exec --trusted
```

Build specific steps:
```bash
drone exec --include=lint
drone exec --include=test
drone exec --include=version,build
drone exec --trusted --include=publish
```

To override and specify custom configuration, edit `custom.env` file and use:
```bash
drone exec --trusted --env-file=custom.env
```

For using secrets locally:
```bash
drone exec --secrets-file=secrets.env
```

