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

