schema "public" {
}

table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = int
  }
  column "name" {
    null = true
    type = varchar(100)
  }
  column "last_name" {
    null = true
    type = varchar(100)
  }
  column "age" {
    null = true
    type = int
  }
  primary_key {
    columns = [column.id]
  }
}

table "blog_posts" {
  schema = schema.public
  column "id" {
    null = false
    type = int
  }
  column "title" {
    null = true
    type = varchar(100)
  }
  column "body" {
    null = true
    type = text
  }
  column "author_id" {
    null = true
    type = int
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "author_fk" {
    columns     = [column.author_id]
    ref_columns = [table.users.column.id]
  }
}
