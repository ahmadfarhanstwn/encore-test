package site

import (
	"encore.dev/storage/sqldb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//encore:service
type Service struct {
	db *gorm.DB
}

var siteDb = sqldb.Named("sites").Stdlib()

func initService() (*Service, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: siteDb,
	}))
	if err != nil {
		return nil, err
	}

	return &Service{db: db}, nil
}
