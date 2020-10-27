#!/bin/bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo $(pwd)
docker build -t away168/hello-metrics:latest .
docker push away168/hello-metrics:latest