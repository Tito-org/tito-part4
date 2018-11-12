#--------------------------#
######  Makefile  ##########
######  Create By  #########
######  Hugo Janasik  ######
######  Intern Developper  #
#--------------------------#

## Methods

help: ## help
	@echo make help : print this display
	@echo "--------------------------------------"
	@echo make clone : clone the github repository
	@echo "--------------------------------------"
	@echo make tito : deploy containers from the docker hub
	@echo "--------------------------------------"
	@echo make tito_stop : stop all the containers from the dockerhub
	@echo "--------------------------------------"
	@echo make tito_rm : remove all the containers from the dockerhub
	@echo "--------------------------------------"
	@echo make tito_rmi : remove all the images from the dockerhub

clone: ## clone github code
	sudo git clone https://github.com/Tito-org/tito-part4.git
	@echo "Clone the github repository Done"

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
