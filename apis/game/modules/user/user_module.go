package user

import (
	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	controller "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/user/controller"
	handler "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/user/handler"
	repository "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/user/repository"
	"github.com/LucasToledoPereira/go-boilerplate/internal/router"
	"gorm.io/gorm"
)

func UserModuleFunc() gbp.ModuleFunc {
	return func(input *gbp.ModuleFuncInput) (err error) {
		migrate(input.Datastore.DB())
		ur := repository.NewUserRepository(input.Datastore)
		uh := handler.NewUserHandler(&ur, input.Filestore)
		uc := controller.NewUserController(uh)
		createRouters(input.Router, uc)
		return nil
	}
}

func createRouters(r *router.Router, c controller.UserController) {
	ig := r.Private.Group("/users")
	ig.GET("info", c.ReadInfo)
	ig.DELETE("", c.DeleteSelf)
	ig.PUT("", c.UpdateSelf)
	ig.POST("upload", c.UploadSelf)
	ig.POST("", c.Create)
	ig.GET("", c.List)
	ig.DELETE(":id", c.Delete)
	ig.PUT(":id", c.Update)
	ig.GET(":id", c.Read)
	ig.POST(":id/upload", c.Upload)
}

func migrate(db *gorm.DB) (err error) {
	if !db.Migrator().HasTable(&entity.User{}) {
		if err = db.Migrator().CreateTable(&entity.User{}); err != nil {
			return err
		}
	}

	hasConstraint := db.Migrator().HasConstraint(&entity.Studio{}, "Users")
	if !hasConstraint {
		if err = db.Migrator().CreateConstraint(&entity.Studio{}, "Users"); err != nil {
			return err
		}
	}

	hasConstraint = db.Migrator().HasConstraint(&entity.Studio{}, "fk_studio_users")
	if !hasConstraint {
		err = db.Migrator().CreateConstraint(&entity.Studio{}, "fk_studio_users")
	}

	hasFileColumn := db.Migrator().HasColumn(&entity.User{}, "filekey")
	if !hasFileColumn {
		err = db.Migrator().AddColumn(&entity.User{}, "filekey")
	}

	return err
}
