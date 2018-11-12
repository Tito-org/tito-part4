######  Makefile  ##########
######  Create By  #########
######  Hugo Janasik  ######
######  Intern Developper  #
#---------------------------#

clone: ## clone github code
	sudo git clone https://github.com/Tito-org/tito-part4.git

build: ## build docker Image
	sudo docker build -t "tito_get_microservice" --build-arg WAR_FILE=./GET_java_garage/vbarbu.war
	sudo docker build -t "tito_book_microservice" BOOK_golang_garage/
	sudo docker build -t "tito_set_microservice" SET_golang_garage/
	sudo docker build -t "tito_reset_microservice" RESET_golang_garage/
	@echo "Build Images Done"

run: ## run docker container
	sudo docker run -d -p 1111:8080 --name="GET_TITO" tito_get_microservice
	sudo docker run -d -p 2222:8080 --name="BOOK_TITO" tito_book_microservice
	sudo docker run -d -p 3333:8080 --name="SET_TITO" tito_set_microservice
	sudo docker run -d -p 5555:8080 --name="RESET_TITO" tito_set_microservice
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
	sudo docker stop RESET_TITO
	@echo "Remove Containers Done"

rmi: ## remove images
	sudo docker rmi tito_get_microservice
	sudo docker rmi tito_book_microservice
	sudo docker rmi tito_set_microservice
	sudo docker rmi tito_reset_microservice
	@echo "Romove Images Done"

.PHONY: clone build run stop rm rmi
