env "development" {
  src = "./db/schema.hcl"
  # db onde sera aplicado as mudancas
  url = "postgres://docker:docker@localhost:5432/docker?search_path=public&sslmode=disable"
  # db temp usado para validar as migracoes e calcular o diff do schema
  dev = "postgres://docker:docker@localhost:5432/postgres?search_path=public&sslmode=disable"
}

env "production" {
  src = "./db/schema.hcl"
  url = "postgres://docker:docker@localhost:5432/docker?search_path=public&sslmode=disable"
}
