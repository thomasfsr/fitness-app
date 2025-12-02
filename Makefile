.PHONY: build run up down

build:
    go build -o api ./cmd/api

run: build
    ./api

up:
    docker-compose up -d --build

down:
    docker-compose down
