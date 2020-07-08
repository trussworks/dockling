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

Try editing the webpage and get that serving in a new docker container