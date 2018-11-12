#--------------------------#
######  Makefile  ##########
######  Create By  #########
######  Hugo Janasik  ######
######  Intern Developper  #
#--------------------------#

help: ## help
	@echo make help : print this display
	@echo "--------------------------------------"
	@echo make clone : clone the github repository
	@echo "--------------------------------------"
	@echo make build : build images from the code
	@echo "--------------------------------------"
	@echo make run : run containers from the docker image that you created
	@echo make tito : deploy containers from the docker hub
	@echo "--------------------------------------"
	@echo make stop : stop all containers from the github
	@echo make tito_stop : stop all the containers from the dockerhub
	@echo "--------------------------------------"
	@echo make rm : remove all containers from the github
	@echo make tito_rm : remove all the containers from the dockerhub
	@echo "--------------------------------------"
	@echo make rmi : remove all images from the github
	@echo make tito_rmi : remove all the images from the dockerhub

clone: ## clone github code
	sudo git clone https://github.com/Tito-org/tito-part4.git
	@echo "Clone the github repository Done"

build: ## build docker Image
	sudo docker build -t "tito_get_microservice" --build-arg WAR_FILE=vbarbu.war GET_java_garage/
	sudo docker build -t "tito_book_microservice" BOOK_golang_garage/
	sudo docker build -t "tito_set_microservice" SET_golang_garage/
	sudo docker build -t "tito_reset_microservice" RESET_golang_garage/
	@echo "Build Images Done"

run: ## run docker container
	sudo docker run -d -p 1111:8080 --name="GET_TITO" tito_get_microservice
	sudo docker run -d -p 2222:8080 --name="BOOK_TITO" tito_book_microservice
	sudo docker run -d -p 3333:8080 --name="SET_TITO" tito_set_microservice
	sudo docker run -d -p 5555:8080 --name="RESET_TITO" tito_reset_microservice
	@echo "Run Containers Done"

stop: ## stop Tito container
	sudo docker stop GET_TITO
	sudo docker stop BOOK_TITO
	sudo docker stop SET_TITO
	sudo docker stop RESET_TITO
	@echo "Stop Containers Done"

rm: ## remove container
	sudo docker rm GET_TITO
	sudo docker rm BOOK_TITO
	sudo docker rm SET_TITO
	sudo docker rm RESET_TITO
	@echo "Remove Containers Done"

rmi: ## remove images
	sudo docker rmi tito_get_microservice
	sudo docker rmi tito_book_microservice
	sudo docker rmi tito_set_microservice
	sudo docker rmi tito_reset_microservice
	@echo "Romove Images Done"

tito: ## deploy container from our docker hub
	sudo docker run -d -p 1111:8080 --name="GET_TITO" hjanasik/get_tito:1.0
	sudo docker run -d -p 2222:8080 --name="BOOK_TITO" hjanasik/book_tito:1.0
	sudo docker run -d -p 3333:8080 --name="SET_TITO" hjanasik/set_tito:1.0
	sudo docker run -d -p 5555:8080 --name="RESET_TITO" hjanasik/reset_tito:1.0
	sudo docker run -d -p 1234:80 --name="Tito_v2" hjanasik/tito:v2
	@echo "Containers deploy from the docker hub Done"

tito_stop: ## stop containers from the dockerhub
	sudo docker stop GET_TITO
	sudo docker stop BOOK_TITO
	sudo docker stop SET_TITO
	sudo docker stop RESET_TITO
	sudo docker stop Tito_v2
	@echo "Stop Tito Containers Done"

tito_rm: ## remove containers from the dockerhub
	sudo docker rm GET_TITO
	sudo docker rm BOOK_TITO
	sudo docker rm SET_TITO
	sudo docker rm RESET_TITO
	sudo docker rm Tito_v2
	@echo "Remove Tito Containers Done"

tito_rmi: ## remove images from the dockerhub
	sudo docker rmi hjanasik/get_tito:1.0
	sudo docker rmi hjanasik/book_tito:1.0
	sudo docker rmi hjanasik/set_tito:1.0
	sudo docker rmi hjanasik/reset_tito:1.0
	sudo docker rmi hjanasik/tito:v2
	@echo "Remove Tito Images Done"


.PHONY: help clone build run stop rm rmi
