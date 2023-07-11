package game

import (
	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	controller "github.com/LucasToledoPereira/go-boilerplate/internal/modules/game/controller"
	handler "github.com/LucasToledoPereira/go-boilerplate/internal/modules/game/handler"
	repository "github.com/LucasToledoPereira/go-boilerplate/internal/modules/game/repository"
	"github.com/LucasToledoPereira/go-boilerplate/internal/router"
	"gorm.io/gorm"
)

func GameModuleFunc() gbp.ModuleFunc {
	return func(input *gbp.ModuleFuncInput) (err error) {
		migrate(input.Datastore.DB())
		gr := repository.NewGameRepository(input.Datastore)
		gh := handler.NewGameHandler(&gr, input.Filestore)
		createRouters(input.Router, controller.NewGameController(gh))
		return nil
	}
}

func createRouters(r *router.Router, c controller.GameController) {
	ig := r.Private.Group("/games")
	ig.GET("", c.List)
	ig.POST("", c.Create)

	tg := ig.Group(":idGame")
	tg.PUT("", c.Update)
	tg.GET("", c.Read)
	tg.DELETE("", c.Delete)
}

func migrate(db *gorm.DB) (err error) {
	if db.Migrator().HasTable(&entity.Game{}) {
		//db.Migrator().DropTable(&entities.Itinerary{})
		return nil
	}
	err = db.Migrator().CreateTable(&entity.Game{})

	hasConstraint := db.Migrator().HasConstraint(&entity.Studio{}, "Games")
	if !hasConstraint {
		db.Migrator().CreateConstraint(&entity.Studio{}, "Games")
	}

	hasConstraint = db.Migrator().HasConstraint(&entity.Studio{}, "fk_studio_games")
	if !hasConstraint {
		db.Migrator().CreateConstraint(&entity.Studio{}, "fk_studio_games")
	}

	return
}
