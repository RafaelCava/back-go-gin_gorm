FROM golang:1.21-alpine

WORKDIR /usr/kitkit-back-go/app

RUN go install github.com/cosmtrek/air@latest

CMD ["tail", "-f", "/dev/null"]