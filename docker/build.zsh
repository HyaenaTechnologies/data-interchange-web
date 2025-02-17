#!/bin/zsh

# Docker Build Package
docker build ./ --file ./data-interchange.Dockerfile --tag data-interchange-web:latest
