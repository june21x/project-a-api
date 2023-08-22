[![Go](https://img.shields.io/badge/Go-v1.20-blue)](https://golang.org/)

# Project A API

Restful API for **Project A**.

The goal of Project A is to create a non-addictive and non-toxic social media where players can explore and interact with the virtual social world in a 2D grid.

_"Project A" is a temporary code name for this project._

## Current Stack

- **API**: Golang with [Gin](https://gin-gonic.com/) framework
- **DB**: [neo4j](https://neo4j.com/)

## Quick Start Guide

### 1. Environment Variables Setup

- **Setting up environment variables needed for development:**

  1.  Run the following scripts. This will create 2 config files (git ignored).

  ```sh
  cp configs/example.config.yaml configs/debug.config.yaml
  cp configs/example.config.yaml configs/release.config.yaml
  ```

  2. Override the values accordingly.

> <span style="color:orange">**Security warning:</span> DO NOT expose these files to the internet!** They contain sensitive info and pose extremely high risk to be abused by malicious parties.

### 2. Install Project Dependencies

```
go get .
```

### 3. Run DB

Use `docker compose` to spin up a neo4j container:

```sh
docker compose -f docker-compose.local-db-only.yml up -d
```

You can browse the DB at `localhost:7474`

### 4. Run API

```sh
go run main.go
```

The API is now run at `localhost:8080`.

You can browse API documentation by Swagger at `localhost:8080/swagger/index.html`

![Alt](https://repobeats.axiom.co/api/embed/ced0edc85cf355645fbebb434a461d988c9a2a31.svg "Repobeats analytics image")
