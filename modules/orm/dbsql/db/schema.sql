-- Create "atlas_schema_revisions" table
CREATE TABLE "atlas_schema_revisions" ("version" character varying NOT NULL, "description" character varying NOT NULL, "type" bigint NOT NULL DEFAULT 2, "applied" bigint NOT NULL DEFAULT 0, "total" bigint NOT NULL DEFAULT 0, "executed_at" timestamptz NOT NULL, "execution_time" bigint NOT NULL, "error" text NULL, "error_stmt" text NULL, "hash" character varying NOT NULL, "partial_hashes" jsonb NULL, "operator_version" character varying NOT NULL, PRIMARY KEY ("version"));
-- Create "person" table
CREATE TABLE "person" ("id" bigserial NOT NULL, "first_name" text NOT NULL, "last_name" text NOT NULL, "email" text NOT NULL, PRIMARY KEY ("id"));
-- Create "place" table
CREATE TABLE "place" ("id" bigserial NOT NULL, "country" text NOT NULL, "city" text NULL, "telcode" integer NOT NULL, PRIMARY KEY ("id"));
-- Create "waypoints" table
CREATE TABLE "waypoints" ("id" bigserial NOT NULL, "name" text NOT NULL, "geom" geometry(point,4326) NOT NULL, PRIMARY KEY ("id"));
-- Create index "waypoints_geom_idx" to table: "waypoints"
CREATE INDEX "waypoints_geom_idx" ON "waypoints" USING gist ("geom");
