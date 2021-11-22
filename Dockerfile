FROM ubuntu:21.10 as postgres

RUN apt-get -y update
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get install -y postgresql

USER root

RUN apt-get install -y golang git

COPY ./ ./

RUN go mod download
RUN go build -o url_shortener ./cmd/main.go

EXPOSE 8080
EXPOSE 5432

USER postgres

RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" &&\
    createdb -O docker docker &&\
    psql -d docker -a -f ./sql/url.sql &&\
    /etc/init.d/postgresql stop

VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

CMD service postgresql start && ./url_shortener -db postgres

FROM ubuntu:21.10 as redis

RUN apt-get -y update
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get install -y redis-server

USER root
RUN apt-get install -y golang git

COPY ./ ./

RUN go mod download
RUN go build -o url_shortener ./cmd/main.go

EXPOSE 8080
EXPOSE 5379

CMD /etc/init.d/redis-server start && ./url_shortener -db redis