#!/usr/bin/env sh

set -e
set -x

cd "$(dirname "${0}")/.."

cleanup() {
   docker rm extract-protobuf-events-generator || true
}
trap cleanup EXIT

# Delete previously generated golang files
find "../../../src/shared/clients/events" -name "*.pb.go" -type f -delete || true

# Protobuf
docker build --progress=plain -f Dockerfile -t protobuf-events-generator ../../../
docker create --name extract-protobuf-events-generator protobuf-events-generator
docker cp extract-protobuf-events-generator:/output/. "../../../src/"
