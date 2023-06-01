-- Create "waypoints" table
CREATE TABLE "waypoints" ("id" bigserial NOT NULL, "name" text NOT NULL, "geom" geometry(point,4326) NOT NULL, PRIMARY KEY ("id"));
