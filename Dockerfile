FROM golang as builder

ENV GO11MODULE=on

WORKDIR /go-lambda

RUN apt-get update && \
    apt-get install -y \
        python3 \
        python3-pip \
        python3-setuptools \
        groff \
        less \
    && pip3 install --upgrade pip \
    && apt-get clean

RUN pip3 --no-cache-dir install --upgrade awscli

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .
