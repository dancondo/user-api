# Users API
> This is a simple users api with authentication module.

- [Users API](#users-api)
  - [Description](#description)
  - [API Docs](#api-docs)
  - [How to Run](#how-to-run)
  - [How to Test](#how-to-test)

## Description

This is a simple golang API that provides a user module

## API Docs

You can generate your swagger by running the command

```bash
$ make swagger
```

After [running the application](#how-to-run) you will find it in the following link:

<http://localhost:4040/docs/swagger/index.html>

## How to Run

You'll need to have [docker](https://docs.docker.com/engine/install/ubuntu/), [docker compose](https://docs.docker.com/compose/cli-command/#install-on-linux) plugin and [golang](https://go.dev/doc/install) installed.

```bash
$ docker-compose up
```

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

### 

```bash
make test/cover
```

### Lint

Just run:

```bash
make lint
```
