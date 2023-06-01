schema "public" {
}

table "waypoints" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "name" {
    null = false
    type = text
  }
  column "geom" {
    null = false
    type = sql("geometry(POINT, 4326)")
  }
  primary_key {
    columns = [
      column.id
    ]
  }
  index "waypoints_geom_idx" {
    type    = GIST
    columns = [column.geom]
  }
}

table "person" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "first_name" {
    null = false
    type = text
  }
  column "last_name" {
    null = false
    type = text
  }
  column "email" {
    null = false
    type = text
  }
  primary_key {
    columns = [
      column.id
    ]
  }
}

table "place" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "country" {
    null = false
    type = text
  }
  column "city" {
    null = true
    type = text
  }
  column "telcode" {
    null = false
    type = int
  }
  primary_key {
    columns = [
      column.id
    ]
  }
}
