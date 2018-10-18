FROM golang

ENV db_type="mysql"
ENV db_username="root"
ENV db_ip="172.18.12.219"
ENV db_password="PASSWORD"
ENV db_name="Test"
ENV tito_ip="172.18.12.219:1234"

RUN apt-get update
RUN apt-get install php -y
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/go-sql-driver/mysql
RUN go build -o binary ./src/check.go ./src/main.go ./src/method.go ./src/reset.go
EXPOSE 8080
WORKDIR /app/src
ENTRYPOINT /app/binary
