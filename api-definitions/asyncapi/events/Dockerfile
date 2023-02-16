# syntax = docker/dockerfile:1.4.1
ARG BUILD_IMAGE=node:18-alpine

FROM $BUILD_IMAGE as build
LABEL maintainer="morteza.alizadeh@gmail.com"

RUN npm install -g @asyncapi/generator

COPY src/utils/go-connectilly-event-template /go-connectilly-event-template
WORKDIR /go-connectilly-event-template
RUN --mount=type=cache,target=~/.yarn YARN_CACHE_FOLDER=~/.yarn yarn install --frozen-lockfile

COPY src/utils/go-connectilly-event-jsonschema-template /go-connectilly-event-jsonschema-template
WORKDIR /go-connectilly-event-jsonschema-template
RUN --mount=type=cache,target=~/.yarn YARN_CACHE_FOLDER=~/.yarn yarn install --frozen-lockfile

COPY api-definitions/asyncapi/events /asyncapi

RUN ag /asyncapi/connectilly/Organization_V1.yaml /go-connectilly-event-template \
    --output /generated-src/connectilly/organization/v1 \
    --force-write \
    -p packageName=organizationv1

# Generate metadata including topic name and json schema

RUN ag /asyncapi/connectilly/Organization_V1.yaml /go-connectilly-event-jsonschema-template \
    --output /generated-src/connectilly/organization/v1/schema \
    --force-write \
    -p packageName=organizationv1
