# Project rowers

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests

```bash
make all build
```

build the application

```bash
make build
```

run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB container

```bash
make docker-down
```

live reload the application

```bash
make watch
```

run the test suite

```bash
make test
```

clean up binary from the last build

```bash
make clean
```

# DATABASE :: Atlas

```sh
# Apply to local database
atlas schema apply --to file://migrations/main.hcl -u sqlite://rowers.db

# Apply to turso database
atlas schema apply -u "$TURSO_DB_URL?authToken=$TURSO_DB_TOKEN" --to sqlite://rowers.db --exclude '_litestream_seq,_litestream_lock'
```
