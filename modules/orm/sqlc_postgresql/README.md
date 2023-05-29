# Setup SLQC

## Install

```
docker compose up -d

go install github.com/go-task/task/v3/cmd/task@v3.25.0

cp env.example .env

task download

task install

task migrate

task seed

task sql

task run
```


## Ref

https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html

https://eltonminetto.dev/en/post/2022-10-22-creating-api-using-go-sqlc/

https://mikemackintosh.com/managing-your-database-migrations-and-seeds-in-go-34d7e0865c43

https://github.com/pressly/goose

https://atlasgo.io/guides/frameworks/sqlc-declarative

https://atlasgo.io/guides/frameworks/sqlc-versioned
