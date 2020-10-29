ID = $(shell docker ps -aqf name=blog)

build:
	docker build --tag blog:0.1 .
run:
	docker run --publish 9090:9090 --detach --name blog blog:0.1
stop:
	docker stop $(ID)
rm:
	docker rm $(ID)
log:
	docker logs $(ID)
