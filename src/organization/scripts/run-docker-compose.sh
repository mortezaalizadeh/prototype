#!/usr/bin/env sh

set -e
set -x

command=${@:-up -d --build}

cd "$(dirname "${0}")/../.."

docker compose -p "connectilly" \
    --profile organization \
    -f docker-compose.yml \
    -f ./organization/docker-compose.yml \
    $command
