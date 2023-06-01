# Setup database/sql

There is a basic connection pool in the database/sql package. There isn’t a lot of ability to control or inspect it, but here are some things you might find useful to know:


* Connection pooling means that executing two consecutive statements on a single database might open two connections and execute them separately. It is fairly common for programmers to be confused as to why their code misbehaves. For example, LOCK TABLES followed by an INSERT can block because the INSERT is on a connection that does not hold the table lock.

* Connections are created when needed and there isn’t a free connection in the pool.

* By default, there’s no limit on the number of connections. If you try to do a lot of things at once, you can create an arbitrary number of connections. This can cause the database to return an error such as “too many connections.”

* In Go 1.1 or newer, you can use db.SetMaxIdleConns(N) to limit the number of idle connections in the pool. This doesn’t limit the pool size, though.

* In Go 1.2.1 or newer, you can use db.SetMaxOpenConns(N) to limit the number of total open connections to the database. Unfortunately, a deadlock bug (fix) prevents db.SetMaxOpenConns(N) from safely being used in 1.2.

* Connections are recycled rather fast. Setting a high number of idle connections with db.SetMaxIdleConns(N) can reduce this churn, and help keep connections around for reuse.


* Keeping a connection idle for a long time can cause problems (like in this issue with MySQL on Microsoft Azure). Try db.SetMaxIdleConns(0) if you get connection timeouts because a connection is idle for too long.

* You can also specify the maximum amount of time a connection may be reused by setting db.SetConnMaxLifetime(duration) since reusing long lived connections may cause network issues. This closes the unused connections lazily i.e. closing expired connection may be deferred.


## Install

```
go install github.com/go-task/task/v3/cmd/task@v3.25.0

curl -sSf https://atlasgo.sh | sh

docker compose up -d

cp env.example .env

task download

task install

task db_migrate

task run
```


## Ref

https://go.dev/doc/database/manage-connections

http://go-database-sql.org/connection-pool.html

http://go-database-sql.org/prepared.html

https://medium.com/propertyfinder-engineering/go-and-mysql-setting-up-connection-pooling-4b778ef8e560

https://koho.dev/understanding-go-and-databases-at-scale-connection-pooling-f301e56fa73

https://entgo.io/docs/sql-integration/#configure-sqldb

https://climbtheladder.com/10-golang-database-best-practices/

https://github.com/jackc/pgx/blob/4ebf1d2e0baa87ab6afbbf6a0947c9fb8a30e9d4/conn_test.go#L384

https://github.com/jackc/pgx/blob/4ebf1d2e0baa87ab6afbbf6a0947c9fb8a30e9d4/copy_from_test.go#L39

https://atlasgo.io/atlas-schema/sql-types#postgresql

https://dev.to/techschoolguru/a-clean-way-to-implement-database-transaction-in-golang-2ba

https://gist.github.com/pseudomuto/0900a7a3605470760579752fcf0fc2b7

https://aloksinhanov.medium.com/query-vs-exec-vs-prepare-in-golang-e7c49212c36c


### Using JSONB

https://www.alexedwards.net/blog/using-postgresql-jsonb


### Using PostGIS with sqlc

https://github.com/kyleconroy/sqlc/issues/231

https://github.com/kyleconroy/sqlc/discussions/384


# Using PostGIS with sqlboiler

https://github.com/volatiletech/sqlboiler/issues/195

https://github.com/theo-m/sqlboiler#types
