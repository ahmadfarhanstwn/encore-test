package site

import "encore.dev/storage/sqldb"

var tododb = sqldb.NewDatabase("sites", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})
