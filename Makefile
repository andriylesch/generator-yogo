.PHONY: docker-run docker-build docker-rm docker

# build docker image
docker-build: 
	docker build -t generator-yogo:latest . 

# run container
docker-run:
	docker run -it --rm -v $(LOCAL_PATH):/home/yeoman -e LOCAL_PATH=$(LOCAL_PATH) --name generator-yogo-container generator-yogo

# build docker image than run container
docker: docker-build docker-run

# remove container and image
docker-rm: 
	docker stop generator-yogo-container && docker rm generator-yogo-container && docker rmi generator-yogo
