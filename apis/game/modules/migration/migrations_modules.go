package migration

import (
	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"gorm.io/gorm"
)

func MigrationModuleFunc() gbp.ModuleFunc {
	return func(input *gbp.ModuleFuncInput) (err error) {
		migrate(input.Datastore.DB())
		return nil
	}
}

func migrate(db *gorm.DB) (err error) {
	if err = createUserRoleType(db); err != nil {
		return
	}
	if err = createAttributeDataType(db); err != nil {
		return
	}
	if err = createFileType(db); err != nil {
		return
	}
	return nil
}

func createUserRoleType(db *gorm.DB) error {
	result := db.Exec("SELECT 1 FROM pg_type WHERE typname = 'user_role'")
	if result.RowsAffected >= 1 {
		return nil
	}
	result = db.Exec("CREATE TYPE user_role as ENUM('OWNER', 'ADMINISTRATOR', 'DEVELOPER', 'LEVEL_DESIGNER', 'CREATOR', 'COMMON')")
	return result.Error
}

func createAttributeDataType(db *gorm.DB) error {
	result := db.Exec("SELECT 1 FROM pg_type WHERE typname = 'attribute_type'")
	if result.RowsAffected >= 1 {
		return nil
	}
	result = db.Exec("CREATE TYPE attribute_type as ENUM('STRING', 'DATE', 'DATETIME', 'NUMBER')")
	return result.Error
}

func createFileType(db *gorm.DB) error {
	result := db.Exec("SELECT 1 FROM pg_type WHERE typname = 'file_type'")
	if result.RowsAffected >= 1 {
		return nil
	}
	result = db.Exec("CREATE TYPE file_type as ENUM('COVER', 'ANY', 'AVATAR')")
	return result.Error
}
