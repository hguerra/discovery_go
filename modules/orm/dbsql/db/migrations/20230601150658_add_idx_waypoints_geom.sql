-- Create index "waypoints_geom_idx" to table: "waypoints"
CREATE INDEX "waypoints_geom_idx" ON "waypoints" USING gist ("geom");
