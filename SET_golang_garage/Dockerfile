FROM golang

ENV db_type="mysql"
ENV db_username="root"
ENV db_password="PASSWORD"
ENV db_ip="172.18.12.219"
ENV db_name="Test"
ENV tito_ip="172.18.12.219:1234"

RUN touch /tmp/.lock
VOLUME /tmp
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/go-sql-driver/mysql
RUN go build -o binary ./src/connect.go ./src/create.go ./src/checkErr.go ./src/getData.go ./src/main.go ./src/method.go ./src/picture.go
EXPOSE 8080
WORKDIR /app/src
ENTRYPOINT /app/binary
