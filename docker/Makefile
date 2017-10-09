.DEFAULT_GOAL := help

docker_server: docker_build docker_up

docker_clean: docker_stop docker_rm

help:
	@echo docker_build:		Build the docker container
	@echo docker_up:		Start the docker container
	@echo docker_stop:		Stop the docker container
	@echo docker_rm:		Remove the docker container
	@echo docker_restart:	Restart the docker container
	@echo docker_ssh: 		Execute an interactive bash shell on the container

docker_build:
	docker-compose build

docker_up:
	docker-compose up

docker_stop:
	docker-compose stop

docker_restart:
	docker-compose restart

docker_rm:
	docker-compose rm

docker_ssh:
	docker exec -it goa-stater /bin/bash