#!/usr/bin/env sh

set -e
set -x

command=${@:-up -d --build}

cd "$(dirname "${0}")/.."

docker compose -p "connectilly" \
    --profile core \
    -f docker-compose.yml \
    -f shared/docker-compose.yml \
    -f organization/docker-compose.yml \
    -f gateway/docker-compose.yml \
    $command
