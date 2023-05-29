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
  # db onde sera aplicado as mudancas
  url = "postgres://docker:docker@localhost:5432/docker?search_path=public&sslmode=disable"
  # db temp usado para validar as migracoes e calcular o diff do schema
  dev = "postgres://docker:docker@localhost:5432/postgres?search_path=public&sslmode=disable"

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
