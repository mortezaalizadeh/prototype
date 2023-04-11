# syntax = docker/dockerfile:1.4.1
ARG BUILD_IMAGE=golang:1.20

FROM $BUILD_IMAGE as build
LABEL maintainer="morteza.alizadeh@gmail.com"

RUN apt-get update -y
RUN apt-get install unzip

ENV GIT_TAG="22.2"
ENV PROTOC_ZIP=protoc-$GIT_TAG-linux-x86_64.zip
RUN curl -OL https://github.com/google/protobuf/releases/download/v$GIT_TAG/$PROTOC_ZIP
RUN unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
RUN unzip -o $PROTOC_ZIP -d /usr/local include/*
RUN rm -f $PROTOC_ZIP

RUN go install github.com/golang/protobuf/protoc-gen-go@latest

RUN apt-get clean

COPY src/go.mod /src/
COPY src/go.sum /src/
COPY src/shared/enterprise /src/shared/enterprise
COPY src/shared/connectillyctl /src/shared/connectillyctl

WORKDIR /src/shared/connectillyctl
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/connectillyctl main.go

RUN mkdir -p /output
COPY api-definitions/protobuf/events /protobuf/events

RUN protoc \
    -I=/protobuf/events/connectilly \
    --go_out=/output \
    Organization_V1.proto

RUN /bin/connectillyctl \
    generate \
    client \
    event \
    --protobufFilePath=/protobuf/events/connectilly/Organization_V1.proto \
    --packageName=v1 \
    --eventType=Event \
    --topicName=organization.v1.event \
    --outputPath=/output/shared/clients/events/connectilly/organization/v1
