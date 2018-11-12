Information Makefile

-----------------

Command:  

make help -> display the help  

make clone -> clone the code from the github  

make build -> build image from the code  

make run -> run the containers from the code  

make tito -> deploy containers from the docker hub  

make stop -> stop the containers  

make tito_stop -> stop the containers from the dockerhub  

make rm -> remove the containers  

make tito_rm -> remove the containers from the dockerhub  

make rmi -> remove the images  

make tito_rmi -> remove the images from the dockerhub  


Information Docker

------------------

GET microservice:   
Image: hjanasik/get_tito:1.0   
Command: docker run -d -p 1111:8080 --name="GET_TITO" hjanasik/get_tito:1.0

------------------

BOOK microservice:   
Image: hjanasik/book_tito:1.0   
Command: docker run -d -p 2222:8080 --name="BOOK_TITO" hjanasik/book_tito:1.0

------------------

SET microservice:   
Image: hjanasik/set_tito:1.0   
Command: docker run -d -p 3333:8080 --name="SET_TITO" hjanasik/set_tito:1.0

------------------

RESET microservice:    
Image: hjanasik/reset_tito:1.0   
Command: docker run -d -p 5555:8080 --name="RESET_TITO" hjanasik/reset_tito:1.0

------------------

Tito_v2 container:   
Image: hjanasik/tito:v2   
Command: docker run -d -p 1234:80 --name="Tito_v2" hjanasik/tito:v2


Enjoy :)

