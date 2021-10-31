# Fiber Gorm Boilerplate
My basic project structure / workflow for creating web application in golang.

Table Of Content
================
- [Requirements](#Requirements)
- [Installations](#Installations)
    - [for development](#for-development)
    - [for development in docker](#for-development-in-docker)
    - [for production](#for-production)
- [Project Structure](#Project-Structure)

Requirements
============

- Go 1.17 
- Postgresql 12
- Docker 20.10 (development only)
- Docker-compose 1.29.2 (development only)

Installations
=============

for development
---------------
1. Copy .env.example to .env
1. Fill .env file based on your working environtment (postgresql port, postgresql password, server port etc)
1. run the program `go run main.go`
1. you can also run unittest using `go test ./...`

for development in docker
-------------------------
1. Copy docker-compose.yml.example to docker-compose.yml
1. Replace every {} in docker-compose.yml based on your need
1. Build and run docker container using `docker-compose up`
1. Copy .env.example to .env
1. Fill .env file based on your docker-compose environtment (postgresql port, postgresql password, server port etc)
1. run the program `go run main.go`
1. you can also run unittest using `go test ./...`

for production
--------------
we assume you use linux server for production
1. Set and export every environtment variable in .env directly on your system
1. Add this environtment variable `export IS_PROD=TRUE`
1. Build the project `go build`
1. run the executable

Project Structure
=================

TODO
