FROM golang:1.21-alpine

RUN mkdir /project

COPY . /project

WORKDIR /project

RUN go build -o app

RUN chmod +x app

ENV HELLO="World v1"

CMD [ "./app" ]