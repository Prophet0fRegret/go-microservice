setup:
#   Build docker image from the application
	docker build -t go-microservice --no-cache --progress=plain .

#   Use Docker-Compose to setup application image created along with a mongo instance
	docker-compose -f mongo-docker-compose.yml up