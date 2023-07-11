package gbp

func newAudit(ds IDatastoreAdapter) error {
	db := ds.DB()
	if db.Migrator().HasTable(&Audit{}) {
		//db.Migrator().DropTable(&entities.Itinerary{})
		return nil
	}

	return db.Migrator().CreateTable(&Audit{})
}
