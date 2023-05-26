# Ent

## Install

```
go run -mod=mod entgo.io/ent/cmd/ent new User Car Group

go generate ./ent

curl -sSf https://atlasgo.sh | sh

# https://atlasgo.io/concepts/url
atlas schema inspect -u "ent://ent/schema" --dev-url "postgresql://docker:docker@localhost:5432/docker?search_path=public&sslmode=disable" --format '{{ sql . "  " }}'

# https://entgo.io/docs/getting-started/#versioned-migrations
atlas migrate diff migration_name \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "postgresql://docker:docker@localhost:5432/docker?search_path=public&sslmode=disable"
```


## Ref

https://entgo.io/docs/getting-started/
