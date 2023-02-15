# syntax = docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine:3
ARG BUILD_IMAGE=golang:1.20

FROM $BUILD_IMAGE as build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/organization/shared \
    mkdir -p /src/organization/organizationctl

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY organization/shared /src/organization/shared
COPY organization/organizationctl /src/organization/organizationctl

WORKDIR /src/organization/organizationctl

RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/organizationctl main.go

FROM build AS test

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM $FINAL_IMAGE as final
LABEL maintainer="morteza.alizadeh@gmail.com"

WORKDIR /connectilly
COPY --from=build /bin/organizationctl /connectilly
COPY --from=build /src/organization/organizationctl/config.yaml /connectilly

RUN /connectilly/organizationctl install-dependecies

# Fix CVE-2022-3996
RUN apk upgrade libssl3 libcrypto3
