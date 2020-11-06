#ID = $(shell docker ps -aqf name=blog)

build:
	docker build --tag graphninja/blog_api .
deploy:
	docker stack deploy --compose-file=docker-compose.yml blog
stop:
	docker stack rm blog
#log:
#	docker logs $(ID)
