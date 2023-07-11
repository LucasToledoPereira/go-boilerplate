package studio

import (
	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	controller "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/studio/controller"
	handler "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/studio/handler"
	repository "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/studio/repository"
	"github.com/LucasToledoPereira/go-boilerplate/internal/router"
	"gorm.io/gorm"
)

func StudioModuleFunc() gbp.ModuleFunc {
	return func(input *gbp.ModuleFuncInput) (err error) {
		migrate(input.Datastore.DB())
		sr := repository.NewStudioRepository(input.Datastore)
		sh := handler.NewStudioHandler(&sr, input.Filestore)
		createRouters(input.Router, controller.NewStudioController(sh))
		return nil
	}
}

func createRouters(r *router.Router, c controller.StudioController) {
	ig := r.Private.Group("/studios")
	ig.POST("upload", c.Upload)
	ig.GET("", c.Read)
	ig.DELETE("", c.Delete)
	ig.PUT("", c.Update)
}

func migrate(db *gorm.DB) (err error) {
	if !db.Migrator().HasTable(&entity.Studio{}) {
		if err = db.Migrator().CreateTable(&entity.Studio{}); err != nil {
			return err
		}
	}

	hasFileColumn := db.Migrator().HasColumn(&entity.Studio{}, "filekey")
	if !hasFileColumn {
		err = db.Migrator().AddColumn(&entity.Studio{}, "filekey")
	}

	return err
}
