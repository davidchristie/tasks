#!/bin/bash

set -ev

docker-compose build

pushd web

REACT_APP_GATEWAY=http://localhost:8080 npm run build

popd
