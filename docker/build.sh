#!/bin/sh

# Docker Build Package
docker build ./ --file ./data-interchange.Dockerfile --tag amd64/data-interchange-server:latest
