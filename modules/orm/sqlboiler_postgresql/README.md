# Setup SQLBoiler with Atlas

## Install

```
go install github.com/go-task/task/v3/cmd/task@v3.25.0

curl -sSf https://atlasgo.sh | sh

docker compose up -d

cp env.example .env

task download

task install

task migrate

task seed

task sql

task run
```


## Ref

https://atlasgo.io/getting-started/

https://atlasgo.io/versioned/diff

https://atlasgo.io/atlas-schema/projects

https://atlasgo.io/lint/analyzers

https://atlasgo.io/declarative/inspect

https://atlasgo.io/guides/migration-tools/goose-import

https://atlasgo.io/guides/frameworks/sqlc-declarative

https://atlasgo.io/guides/frameworks/sqlc-versioned

https://thedevelopercafe.com/articles/sql-in-go-with-sqlboiler-ac8efc4c5cb8

https://gh.atlasgo.cloud/explore

https://koho.dev/understanding-go-and-databases-at-scale-connection-pooling-f301e56fa73
