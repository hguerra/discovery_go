-- Create "atlas_schema_revisions" table
CREATE TABLE "atlas_schema_revisions" ("version" character varying NOT NULL, "description" character varying NOT NULL, "type" bigint NOT NULL DEFAULT 2, "applied" bigint NOT NULL DEFAULT 0, "total" bigint NOT NULL DEFAULT 0, "executed_at" timestamptz NOT NULL, "execution_time" bigint NOT NULL, "error" text NULL, "error_stmt" text NULL, "hash" character varying NOT NULL, "partial_hashes" jsonb NULL, "operator_version" character varying NOT NULL, PRIMARY KEY ("version"));
-- Create "users" table
CREATE TABLE "users" ("id" bigserial NOT NULL, "name" character varying(100) NULL, "last_name" character varying(100) NULL, "age" integer NULL, PRIMARY KEY ("id"));
-- Create "blog_posts" table
CREATE TABLE "blog_posts" ("id" bigserial NOT NULL, "title" character varying(100) NULL, "body" text NULL, "author_id" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "author_fk" FOREIGN KEY ("author_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
