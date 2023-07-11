package collection

import (
	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	controller "github.com/LucasToledoPereira/go-boilerplate/internal/modules/collection/controller"
	handler "github.com/LucasToledoPereira/go-boilerplate/internal/modules/collection/handler"
	repository "github.com/LucasToledoPereira/go-boilerplate/internal/modules/collection/repository"
	"github.com/LucasToledoPereira/go-boilerplate/internal/router"
	"gorm.io/gorm"
)

func CollectionModuleFunc() gbp.ModuleFunc {
	return func(input *gbp.ModuleFuncInput) (err error) {
		migrate(input.Datastore.DB())
		cr := repository.NewCollectionRepository(input.Datastore)
		ch := handler.NewCollectionHandler(&cr, input.Filestore)
		cc := controller.NewCollectionController(ch)
		createRouters(input.Router, cc)
		return nil
	}
}

func createRouters(r *router.Router, c controller.CollectionController) {
	tg := r.Private.Group("/games/:idGame")
	tr := tg.Group("collections")
	tr.GET("", c.List)
	tr.POST("", c.Create)

	tri := tr.Group(":idCollection")
	tri.PUT("", c.Update)
	tri.GET("", c.Read)
	tri.DELETE("", c.Delete)
	tri.POST("avatar", c.UploadAvatar)
	tri.POST("cover", c.UploadCover)

	ta := tri.Group("attributes")
	ta.GET("", c.ListAttributes)
	ta.POST("", c.CreateAttribute)
	ta.PUT(":idAttribute", c.UpdateAttribute)
	ta.GET(":idAttribute", c.ReadAttribute)
	ta.DELETE(":idAttribute", c.DeleteAttribute)
}

func migrate(db *gorm.DB) (err error) {
	if !db.Migrator().HasTable(&entity.Collection{}) {
		err = db.Migrator().CreateTable(&entity.Collection{})
		if err != nil {
			return err
		}
	}

	hasConstraint := db.Migrator().HasConstraint(&entity.Game{}, "Collections")
	if !hasConstraint {
		if err = db.Migrator().CreateConstraint(&entity.Game{}, "Collections"); err != nil {
			return err
		}
	}

	hasConstraint = db.Migrator().HasConstraint(&entity.Game{}, "fk_game_collections")
	if !hasConstraint {
		if err = db.Migrator().CreateConstraint(&entity.Game{}, "fk_game_collections"); err != nil {
			return err
		}
	}

	if err = migrateCollectionAttributes(db); err != nil {
		return err
	}

	return nil
}

func migrateCollectionAttributes(db *gorm.DB) (err error) {
	if !db.Migrator().HasTable(&entity.CollectionAttributes{}) {
		err = db.Migrator().CreateTable(&entity.CollectionAttributes{})
		if err != nil {
			return err
		}
	}

	hasConstraint := db.Migrator().HasConstraint(&entity.Collection{}, "Attributes")
	if !hasConstraint {
		if err = db.Migrator().CreateConstraint(&entity.Collection{}, "Attributes"); err != nil {
			return err
		}
	}

	hasConstraint = db.Migrator().HasConstraint(&entity.Collection{}, "fk_collection_attributes")
	if !hasConstraint {
		if err = db.Migrator().CreateConstraint(&entity.Collection{}, "fk_collection_attributes"); err != nil {
			return err
		}
	}
	return nil
}
