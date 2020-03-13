package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDB() (*gorm.DB, error) {
	DBMS := "postgres"
	pgConfig := "host=localhost port=5432 user=postgres dbname=sample_db password=postgres"

	return gorm.Open(DBMS, pgConfig)
}
