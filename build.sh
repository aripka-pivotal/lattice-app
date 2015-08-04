#!/usr/bin/env bash

echo "Compiling for linux..."
GOOS=linux GOARCH=amd64 go build .

echo "Constructing Dockerimage"
docker build -t="rippmn/lattice-app" .

echo "Cleaning up..."
rm lattice-app
