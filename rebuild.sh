#!/bin/zsh
docker compose down -v
docker container prune
docker image prune
docker image rm dockerlearn_web
docker compose up --build