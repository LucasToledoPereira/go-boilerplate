package gameapi

import (
	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/auth"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/collection"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/game"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/migration"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/studio"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/template"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/user"
)

func LoadGameAPI(builder *gbp.Builder) {
	builder.Register(migration.MigrationModuleFunc(), auth.AuthModuleFunc())
	builder.Register(studio.StudioModuleFunc())
	builder.Register(user.UserModuleFunc())
	builder.Register(game.GameModuleFunc())
	builder.Register(template.TemplateModuleFunc())
	builder.Register(collection.CollectionModuleFunc())
}
