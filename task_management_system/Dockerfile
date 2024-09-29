# syntax=docker/dockerfile:1

FROM golang:latest

RUN mkdir /temp

WORKDIR /temp
ADD . /temp/

RUN go build -o ./out/task_management_system .

EXPOSE 3000

ENTRYPOINT ["./out/task_management_system"]