# Re Tech Challenge — Packaging API

A Go service that exposes a Packaging API backed by MongoDB. This repository includes a Makefile, a Dockerfile, and a docker-compose setup to make development and execution straightforward.

## TL;DR
- The API can be accessed live here: https://urchin-app-qzpip.ondigitalocean.app/
- Create a `.env` file based on `.env.example` to run the project locally
- Use `make docker-run` to run the API locally
- API default port: 3000
- Use `curl http://localhost:3000/health` to check if the API is running

## Available endpoints
- `POST /packaging` - Creates a new package of the specified `size`
- `GET /packaging` - Gets all available packages
- `GET /packaging/amount/{amount}` - Gets the best package combo for the specified amount following the rules:
  - Rule 1: Only whole packs can be sent. Packs cannot be broken open.
  - Rule 2: Within the constraints of Rule 1 above, send out the least amount of items to fulfil the order.
  - Rule 3: Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each
    order.
- `DELETE /packaging/{id}` - Deletes the package with the specified ID.

## Overview
- Database: MongoDB (containerized)
- Container orchestration for local dev: docker-compose

## Prerequisites
- Docker (20.10+) and Docker Compose v2 (typically available as `docker compose`)
- Make (available by default on macOS/Linux; on Windows, use Git Bash or WSL, or run commands directly)
- Optional for local (non-Docker) run: Go 1.24+ installed locally


## Quick start (recommended): Run with Docker Compose
This is the simplest way to run both the API and MongoDB with persistence.

- Create the `.env` file
  - `cp .env.example .env`

- Start API and MongoDB (builds images as needed):
  - Using Makefile: `make docker-run`
  - Or directly: `docker compose up -d --build api`

- Check running services:
  - `docker compose ps`

- Tail logs:
  - API only: `docker compose logs -f api`
  - Mongo only: `docker compose logs -f mongo`
  - Everything: `docker compose logs -f`

- Stop services (keep volumes/data):
  - `docker compose down`

- Remove everything including persistent data:
  - `docker compose down -v`

## Environment variables
Environment variables are defined in `.env.example` and loaded by docker-compose and Makefile targets.

Key variables:
- SV_PORT: API port (default 3000)
- SV_SERVICE_NAME: service name (default packaging-api)
- SV_ENV: environment (default local)
- LOG_LEVEL: log level (info, debug, warn, error)
- LOG_JSON: output logs in JSON (true/false)
- DB_URL: MongoDB connection string (set to internal docker-compose Mongo URL)

Important: The compose setup uses an internal Mongo service named `mongo`. `.env` is preconfigured with:
```
DB_URL=mongodb://admin:mongopw@mongo:27017/?authSource=admin
```

### MongoDB data persistence
The MongoDB container stores its data in a named volume `mongo_data`. This keeps your data across restarts of the compose stack.

## Using the Makefile
The Makefile provides convenient shortcuts. Common targets:

- `make db-local`
  - Starts only MongoDB via docker compose (detached). Useful for local runs without containerizing the API.

- `make run-api`
  - Ensures dependencies (`go mod tidy`) and vendor folder, then runs the API locally using your machine’s Go toolchain. Requires `DB_URL` to point at a reachable MongoDB. With `db-local` running, the default `.env` works.

- `make build`
  - Builds the project into `./build` with vendored dependencies and copies `.env` and Swagger spec into `./build/static` (if present). Cross-platform directory creation is handled.

- `make docker-build`
  - Builds the Docker image using the Dockerfile. An optional `VERSION_IMAGE` may be used to tag the image (see Makefile for logic); by default, it tags `packaging-api`.

- `make docker-run`
  - Runs the API through docker compose (`docker compose up -d --build api`). This also starts MongoDB.

- `make test`
  - Runs all tests.

- `make test-cover`
  - Runs tests with coverage and opens an HTML coverage report.

## Running without Docker (local Go runtime)
1. Start MongoDB in Docker (detached):
   - `make db-local`
2. Run the API against that Mongo instance:
   - `make run-api`
3. The API will be available at `http://localhost:3000` by default.

If you already have a local MongoDB, adjust `DB_URL` in `.env` accordingly (e.g., `mongodb://localhost:27017`).

## Dockerfile details
- Multi-stage build: Go builder image compiles the binaries into `/go/re-tech-challenge/build`.
- Runtime image: `alpine:latest` with minimal utilities.
- The compiled binaries are copied to `/usr/local/bin` and the container runs the `api` binary by default.
- Exposes port defined by `SV_PORT` (defaults to 3000). The compose file maps this to your host.

## docker-compose services
- mongo
  - Image: `mongo:6`
  - Credentials: `admin` / `mongopw`
  - Port mapping: `27017:27017`
  - Volume: `mongo_data:/data/db` (persistent)

- api
  - Built from the local Dockerfile
  - Loads environment from `.env`
  - Port mapping: `3000:3000`
  - Depends on `mongo`

## Verifying the API routing for available endpoints. If Swagger is enabled, look for `/swagger` or related routes.
- Example curl (replace with an actual endpoint from the project):
  - `curl http://localhost:3000/health`

## Health endpoint
- GET /health
  - Returns 200 OK with body: "Ok".
  - Local (compose): http://localhost:3000/health
