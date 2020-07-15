Dockling
========

This is a small repository built for getting familiar with Docker. 

There is a go program called dockling that prints out a bunch of runtime stats on launch and then serves a simple "hello" webpage. 


Prereqs:

1. Install docker: https://hub.docker.com/editions/community/docker-ce-desktop-mac/
2. Install go: `brew install go`
2. Clone this repo: `git clone github.com/trussworks/dockling`


### Exercise 1

1. Build the docker _image_ specified in `Dockerfile`: `docker build -t dockling .`
2. In side by side windows run dockling locally and in a docker _container_:
	1. Locally: `go run ./cmd/dockling`
	2. Docker: `docker run -p 8043:8042 dockling`

Compare the logs from running the two.

Hit the two web servers at http://localhost:8042 (local) and http://localhost:8043 (docker)

Look at the output of `docker ps` and `docker images`

Try editing the webpage and get that serving in a new docker _container_

### Exercise 2

Use the little website at /name_saver. It requires redis to be available. There are official redis images available here: https://hub.docker.com/_/redis

1. Get a redis image running, (make sure you expose a port!) and try and get the dockling server to connect to it locally
* get dockling running (`go run ./cmd/dockling`), visit http://localhost:8042/name_saver
* it will tell you wether it can see redis or not. Once it does, it will use redis to save a list of names. 

2. Get the dockling server to connect to it while it is _also_ running in a docker container
* you will need to create your own bridge network to make the two containers be able to see eachother by hostname: https://docs.docker.com/network/bridge/ (the default bridge network does not route via hostname)
* `docker network create dockling-net`
* you can configure `dockling` with the environment variables REDIS_HOST and REDIS_PORT to try and connect to redis-in-docker
* when they are both running, you should be able to visit http://localhost:8043/name_server and save names like before

Here are two base64 encoded invocations that worked for me to run dockling and redis in containers and talk to eachother:

ZG9ja2VyIHJ1biAtcCA4MDQzOjgwNDIgLWUgUkVESVNfSE9TVD1kb2NrbGluZy1yZWRpcyAtLW5ldCBkb2NrbGluZy1uZXQgZG9ja2xpbmc=

ZG9ja2VyIHJ1biAtLW5hbWUgZG9ja2xpbmctcmVkaXMgLXAgNjM3OTo2Mzc5IC0tbmV0IGRvY2tsaW5nLW5ldCByZWRpcw==

There's a couple ways to do this, you can either keep redis running and use `docker network disconnect/connect` to attach it to dockling-net or you can start a new container on dockling-net with `docker run`
