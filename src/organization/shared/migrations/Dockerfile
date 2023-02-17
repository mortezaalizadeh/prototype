# syntax = docker/dockerfile:1.4.1
ARG BUILD_IMAGE=microbusiness/atlas:latest

FROM $BUILD_IMAGE as test
LABEL maintainer="morteza.alizadeh@gmail.com"

WORKDIR /migrations
COPY organization/shared/migrations /migrations

CMD /atlas migrate apply --dir file:///migrations/ --url $CONNECTILLY_POSTGRES_CONNECTIONSTRING
