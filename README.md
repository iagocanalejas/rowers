# Project rowers

One Paragraph of project description goes here

## Technologies

### MakeFile

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

live reload the application

```bash
make watch
```

clean up binary from the last build

```bash
make clean
```

### [ECHO](https://echo.labstack.com/)

### [Templ](https://templ.guide/)

### [HTMX](https://htmx.org/)

### [TailwindCss](https://tailwindcss.com/)

### [Turso Database (Sqlite)](https://turso.tech/)

### [Atlas](https://atlasgo.io/)

```sh
# Apply to local database
atlas schema apply --to file://migrations/main.hcl -u sqlite://rowers.db

# Apply to turso database
atlas schema apply -u "$TURSO_DB_URL?authToken=$TURSO_DB_TOKEN" --to sqlite://rowers.db --exclude '_litestream_seq,_litestream_lock'
```
