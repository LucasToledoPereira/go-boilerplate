package audit

import "github.com/LucasToledoPereira/go-boilerplate/adapters/datastore"

func New(ds datastore.IDatastoreAdapter) error {
	db := ds.DB()
	if db.Migrator().HasTable(&Audit{}) {
		//db.Migrator().DropTable(&entities.Itinerary{})
		return nil
	}

	return db.Migrator().CreateTable(&Audit{})
}
