package examplestdlib

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"os"

	geom "github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkb"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
	"github.com/twpayne/go-geom/encoding/geojson"

	"github.com/hguerra/discovery_go/modules/orm/dbsql/internal/db"
	"github.com/hguerra/discovery_go/modules/orm/dbsql/internal/util"
)

// A Waypoint is a location with an identifier and a name.
type Waypoint struct {
	ID       int             `json:"id"`
	Name     string          `json:"name"`
	Geometry json.RawMessage `json:"geometry"`
}

func helloWorld(db *sql.DB) error {
	var greeting string
	err := db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return err
	}
	log.Println(greeting)
	return nil
}

// populateDB demonstrates populating a PostgreSQL/PostGIS database using
// ref: https://github.com/twpayne/go-geom/blob/master/examples/postgis/main.go#LL52C1-L52C36
func populateDB(ctx context.Context, db *sql.DB) error {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO public.waypoints ("name", "geom") VALUES($1, $2)`)
	if err != nil {
		return err
	}

	for _, waypoint := range []struct {
		name string
		geom *geom.Point
	}{
		{"São José dos Campos - SP", geom.NewPoint(geom.XY).MustSetCoords([]float64{-45.88694, -23.17944}).SetSRID(4326)},
	} {
		ewkbhexGeom, err := ewkbhex.Encode(waypoint.geom, ewkbhex.NDR)
		if err != nil {
			return err
		}
		if _, err := stmt.ExecContext(ctx, waypoint.name, ewkbhexGeom); err != nil {
			return err
		}
	}

	err = tx.Commit()
	stmt.Close()
	return err
}

// writeGeoJSON demonstrates reading data from a database in EWKB format and
// writing it as GeoJSON.
// ref https://github.com/twpayne/go-geom/blob/master/examples/postgis/main.go
func writeGeoJSON(ctx context.Context, db *sql.DB, w io.Writer) error {
	rows, err := db.QueryContext(
		ctx,
		`SELECT id, name, ST_AsEWKB(geom) FROM waypoints ORDER BY id ASC`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var ewkbPoint ewkb.Point
		if err := rows.Scan(&id, &name, &ewkbPoint); err != nil {
			return err
		}
		geometry, err := geojson.Marshal(ewkbPoint.Point)
		if err != nil {
			return err
		}
		if err := json.NewEncoder(w).Encode(&Waypoint{
			ID:       id,
			Name:     name,
			Geometry: geometry,
		}); err != nil {
			return err
		}
	}
	return nil
}

func Run() {
	ctx := context.Background()

	db, err := db.OpenDB(ctx, os.Getenv("DATABASE_URL"), true)
	util.Catch(err)
	defer db.Close()

	util.Catch(helloWorld(db))
	util.Catch(populateDB(ctx, db))
	util.Catch(writeGeoJSON(ctx, db, os.Stdout))
}
