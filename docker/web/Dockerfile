FROM golang:1.19.3

RUN apt update && apt install -y git
RUN go install github.com/cosmtrek/air@latest

COPY ../../web /var/golang/web


WORKDIR /var/golang/web
RUN go mod download

EXPOSE 8000

CMD [ "go", "run" , "main.go"]
