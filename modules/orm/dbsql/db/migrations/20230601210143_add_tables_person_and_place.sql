-- Create "person" table
CREATE TABLE "person" ("id" bigserial NOT NULL, "first_name" text NOT NULL, "last_name" text NOT NULL, "email" text NOT NULL, PRIMARY KEY ("id"));
-- Create "place" table
CREATE TABLE "place" ("id" bigserial NOT NULL, "country" text NOT NULL, "city" text NULL, "telcode" integer NOT NULL, PRIMARY KEY ("id"));
