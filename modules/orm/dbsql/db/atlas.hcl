lint {
  git {
    base = "main"
  }

  destructive {
    error = true
  }

  data_depend {
    error = true
  }

  incompatible {
    error = true
  }

  concurrent_index {
    error        = true
    check_create = true
    check_drop   = true
    check_txmode = true
  }

  naming {
    error   = true
    match   = "^[a-z]+$"
    message = "must be lowercase"

    index {
      match   = "^[a-z]+_idx$"
      message = "must be lowercase and end with _idx"
    }
  }
}

env "development" {
  src = "file://db/schema.hcl"
  url = "postgres://docker:docker@localhost:5432/docker?search_path=public&sslmode=disable"
  dev = "postgres://docker:docker@localhost:5432/template_postgis?search_path=public&sslmode=disable"

  migration {
    dir = "file://db/migrations"
  }
}

env "production" {
  src = "./db/schema.hcl"
  url = "postgres://docker:docker@localhost:5432/docker?search_path=public&sslmode=disable"

  migration {
    dir = "file://db/migrations"
  }
}
