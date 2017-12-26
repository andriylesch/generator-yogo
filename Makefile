.PHONY: docker-run docker-build docker-rm docker

# build docker image
docker-build: 
	docker build -t generator-golang:latest . 

# run container
docker-run:
	docker run -it -v $(LOCAL_PATH):/home/yeoman --name generator-golang-container generator-golang

# remove container and image
docker-rm: 
	docker stop generator-golang-container && docker rm generator-golang-container && docker rmi generator-golang
