FROM golang:latest

LABEL maintainer="Ana <anajulia.negrixavier@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod downloud
 
COPY .  . 

ENV PORT 9000 

RUN go build
CMD ["./novo_projeto"]
