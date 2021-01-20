#!/bin/bash

echo "Generating backend code..."
go generate ./...

echo "Generating frontend code..."
pushd web
npm run generate
popd
