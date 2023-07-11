package template

import (
	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	controller "github.com/LucasToledoPereira/go-boilerplate/internal/modules/template/controller"
	handler "github.com/LucasToledoPereira/go-boilerplate/internal/modules/template/handler"
	repository "github.com/LucasToledoPereira/go-boilerplate/internal/modules/template/repository"
	"github.com/LucasToledoPereira/go-boilerplate/internal/router"
	"gorm.io/gorm"
)

func TemplateModuleFunc() gbp.ModuleFunc {
	return func(input *gbp.ModuleFuncInput) (err error) {
		migrate(input.Datastore.DB())
		tr := repository.NewTemplateRepository(input.Datastore)
		th := handler.NewTemplateHandler(&tr, input.Filestore)
		tc := controller.NewTemplateController(th)
		createRouters(input.Router, tc)
		return nil
	}
}

func createRouters(r *router.Router, c controller.TemplateController) {
	tg := r.Private.Group("/games/:idGame")
	tr := tg.Group("templates")
	tr.GET("", c.List)
	tr.POST("", c.Create)

	tri := tr.Group(":idTemplate")

	tri.PUT("", c.Update)
	tri.GET("", c.Read)
	tri.DELETE("", c.Delete)
	tri.POST("image", c.UploadImage)
	tri.POST("files", c.UploadFiles)
	tri.GET("files", c.ListFiles)
	tri.DELETE("files", c.DeleteFiles)

	ta := tri.Group("attributes")
	ta.GET("", c.ListAttributes)
	ta.POST("", c.CreateAttribute)
	ta.PUT(":idAttribute", c.UpdateAttribute)
	ta.GET(":idAttribute", c.ReadAttribute)
	ta.DELETE(":idAttribute", c.DeleteAttribute)
}

func migrate(db *gorm.DB) (err error) {
	if !db.Migrator().HasTable(&entity.Template{}) {
		err = db.Migrator().CreateTable(&entity.Template{})
		if err != nil {
			return err
		}
	}

	hasConstraint := db.Migrator().HasConstraint(&entity.Game{}, "Templates")
	if !hasConstraint {
		if err = db.Migrator().CreateConstraint(&entity.Game{}, "Templates"); err != nil {
			return err
		}
	}

	hasConstraint = db.Migrator().HasConstraint(&entity.Game{}, "fk_game_templates")
	if !hasConstraint {
		if err = db.Migrator().CreateConstraint(&entity.Game{}, "fk_game_templates"); err != nil {
			return err
		}
	}

	if err = migrateTemplateAttributes(db); err != nil {
		return err
	}

	if err = migrateTemplateFiles(db); err != nil {
		return err
	}

	return nil
}

func migrateTemplateAttributes(db *gorm.DB) (err error) {
	if !db.Migrator().HasTable(&entity.TemplateAttributes{}) {
		err = db.Migrator().CreateTable(&entity.TemplateAttributes{})
		if err != nil {
			return err
		}
	}

	hasConstraint := db.Migrator().HasConstraint(&entity.Template{}, "Attributes")
	if !hasConstraint {
		if err = db.Migrator().CreateConstraint(&entity.Template{}, "Attributes"); err != nil {
			return err
		}
	}

	hasConstraint = db.Migrator().HasConstraint(&entity.Template{}, "fk_template_attributes")
	if !hasConstraint {
		if err = db.Migrator().CreateConstraint(&entity.Template{}, "fk_template_attributes"); err != nil {
			return err
		}
	}
	return nil
}

func migrateTemplateFiles(db *gorm.DB) (err error) {
	if !db.Migrator().HasTable(&entity.TemplateFiles{}) {
		err = db.Migrator().CreateTable(&entity.TemplateFiles{})
		if err != nil {
			return err
		}
	}

	hasConstraint := db.Migrator().HasConstraint(&entity.Template{}, "Files")
	if !hasConstraint {
		if err = db.Migrator().CreateConstraint(&entity.Template{}, "Files"); err != nil {
			return err
		}
	}

	hasConstraint = db.Migrator().HasConstraint(&entity.Template{}, "fk_template_files")
	if !hasConstraint {
		if err = db.Migrator().CreateConstraint(&entity.Template{}, "fk_template_files"); err != nil {
			return err
		}
	}
	return nil
}
