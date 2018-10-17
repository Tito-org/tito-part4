ALEA Microservice:


  Code: Golang
  Docker Image: hjanasik/alea_golang_garage:1.0
  Docker Port: 8080
  
  Command:
    docker build -t name . = build a docker image
    docker run -d -p PORT:8080 + Docker Image
    
  Instruction Code:
    http://IP:PORT/go = launch the randomizer
    http://IP:PORT/stop = stop the randomizer


RESET Microservice:


  Code: Golang
  Docker Image: hjanasik/reset_golang_garage:1.0
  Docker Port: 8080
  
  Command:
    docker build -t name . = build a docker image
    docker run -d -p PORT:8080 + Docker Image
  
  Instruction Code:
    http://IP:PORT = Reset the status of each car
    
    
SET Microservice:


  Code: Golang
  Docker Image: hjanasik/set_garage_golang:1.0
  Docker Port: 8080
  
  Command:
    docker build -t name . = build a docker image
    docker run -d -p PORT:8080 + Docker Image
  
  Instruction Code:
    http://IP:PORT = Display a HTML page where you can add element in DB
    
    
BOOK Microservice


  Code: Golang
  Docker Image: hjanasik/book_garage_golang:1.1
  Docker Port: 8080
  
  Command:
    docker build -t name . = build a docker image
    docker run -d -p PORT:8080 + Docker Image
  
  Instruction Code:
    http://IP:PORT = Display HTML page where you can book a car from our DB
