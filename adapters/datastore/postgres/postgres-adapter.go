package postgresadapter

import (
	"fmt"

	"github.com/LucasToledoPereira/go-boilerplate/adapters/datastore"
	"github.com/LucasToledoPereira/go-boilerplate/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresAdapter struct {
	data *datastore.Data
}

func (pa *PostgresAdapter) New() (err error) {
	db, err := connectDB()

	if err != nil {
		return err
	}

	pa.data = &datastore.Data{
		DB: db,
	}

	return nil
}

func (pa *PostgresAdapter) Migrate() (err error) {
	return nil
}

func (pa *PostgresAdapter) DB() (db *gorm.DB) {
	return pa.data.DB
}

func connectDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.C.Database.Host,
		config.C.Database.Port,
		config.C.Database.User,
		config.C.Database.Password,
		config.C.Database.DBName,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
