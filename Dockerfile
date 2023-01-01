FROM golang:latest AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

RUN apt-get update -y &&\
    apt-get upgrade -y &&\
    curl -sL https://deb.nodesource.com/setup_16.x | bash - &&\
    apt-get install nodejs -y &&\
    npm install -g yarn

COPY ./ /root/workdir
WORKDIR /root/workdir
RUN sh devtools.sh
