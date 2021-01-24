#!/bin/bash

set -ev

pushd test/acceptance

docker-compose down -v
docker-compose up -d --force-recreate

pushd cypress

npm install
npm test

popd

popd
