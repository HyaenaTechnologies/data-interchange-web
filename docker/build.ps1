#!bin/pwsh

# Docker Build Package
docker build ./ --file ./data-interchange.Dockerfile --tag data-interchange-server:latest
