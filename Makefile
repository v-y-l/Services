NAME=service
TAG=vyl/blog-${NAME}
VER=v1.0

all: clean build run

build:
	go get github.com/gorilla/mux
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o service
	docker build -t $(TAG) -t $(TAG):$(VER) .

run:
	# Expose and publish on the same port
	# https://nickjanetakis.com/blog/docker-tip-59-difference-between-exposing-and-publishing-ports
	docker run -d -p 80:80 -e PORT=80 --name=$(NAME) $(TAG)

clean:
	-docker stop $(NAME)
	-docker rm $(NAME)

push:
	docker push $(TAG)
	docker push $(TAG):$(VER)
