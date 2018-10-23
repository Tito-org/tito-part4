FROM golang

ENV db_type="mysql"
ENV db_username="root"
ENV db_password="PASSWORD"
ENV db_ip="172.18.12.219"
ENV db_name="Test"

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/go-sql-driver/mysql
RUN go build -o binary ./src/main.go ./src/reset.go ./src/method.go
EXPOSE 8080
WORKDIR /app/src
ENTRYPOINT /app/binary
