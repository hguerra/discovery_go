-- Create "users" table
CREATE TABLE "users" ("id" bigserial NOT NULL, "name" character varying(100) NULL, "last_name" character varying(100) NULL, "age" integer NULL, PRIMARY KEY ("id"));
-- Create "blog_posts" table
CREATE TABLE "blog_posts" ("id" bigserial NOT NULL, "title" character varying(100) NULL, "body" text NULL, "author_id" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "author_fk" FOREIGN KEY ("author_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
