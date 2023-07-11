package datastore

import "gorm.io/gorm"

type Data struct {
	DB *gorm.DB
}

type IDatastoreAdapter interface {
	New() (err error)
	Migrate() (err error)
	DB() (db *gorm.DB)
}
