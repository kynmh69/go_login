FROM golang:1.19.3

RUN apt update && apt install -y git

COPY ../../web /var/golang/web


WORKDIR /var/golang/web
RUN go mod download

EXPOSE 8000
