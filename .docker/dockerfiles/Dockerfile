FROM golang:1.21

WORKDIR /usr/kitkit-back-go/app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

CMD ["tail", "-f", "/dev/null"]