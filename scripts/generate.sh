#!/bin/bash

set -ev

go generate ./...

pushd web

npm run generate

popd
