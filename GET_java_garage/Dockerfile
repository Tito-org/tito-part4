FROM openjdk:8-jdk-alpine

ENV db_type="mysql"
ENV db_ip="172.18.12.219"
ENV db_username="root"
ENV db_password="PASSWORD"
ENV db_connector="jdbc"
ENV db_name="Test"
ENV tito_ip="172.18.12.219:1234"

RUN touch /tmp/.lock
VOLUME /tmp
ARG WAR_FILE
COPY ${WAR_FILE} app.war
EXPOSE 8080
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/app.war"]

