# Users API
> This is a simple users api with authentication module.

- [Users API](#users-api)
  - [Description](#description)
  - [Folder Structure](#folder-structure)
  - [API Docs](#api-docs)
  - [How to Run](#how-to-run)
  - [How to Test](#how-to-test)

## Description

This is a simple golang API that provides a user module

## Folder Structure

In this project you can find the following folders

- [api](api)
    - The api module.
    - In this folder you can find the API setup as well as the declaration of the endpoints.
- [cmd](cmd)
    - The application commands
    - Each command has a file to start the desired part of the application.
- [common](common)
    - The application shared logic
- [docs](docs)
    - The application documentation.
- [pkg](pkg)
    - Standalone packages that can be reused in other applications.
- [repository](repository)
    - The application repositories, that may include databases and cache.


## API Docs

Install swagger with:

```bash
$ go get -u github.com/swaggo/swag/cmd/swag
```

You can generate your swagger by running the command:

```bash
$ make swagger
```

After [running the application](#how-to-run) you will find it in the following link:

<http://localhost:4040/docs/swagger/index.html>

## How to Run

### With Docker Compose

You'll need to have [docker](https://docs.docker.com/engine/install/ubuntu/), [docker compose](https://docs.docker.com/compose/cli-command/#install-on-linux) plugin and [golang](https://go.dev/doc/install) installed.

* **To run for the first time or after changes**

```bash
$ docker-compose up --build
```

* **To run if no changes were made**

```bash
$ docker-compose up
```

### With local dependencies

* **Install dependencies**

```bash
$ make install
```

* **Running local dependencies** if you want to run the application using local dependencies with docker, run the command below

```bash
$ make deps/up
```

* **Run api** in dev mode

```bash
$ make run/api
```

## How to Test

```bash
make test/cover
```
