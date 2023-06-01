package main

import (
	"github.com/hguerra/discovery_go/modules/orm/dbsql/internal/examplesqlx"
	"github.com/hguerra/discovery_go/modules/orm/dbsql/internal/examplestdlib"
)

func main() {
	examplestdlib.Run()
	examplesqlx.Run()
}
