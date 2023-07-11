package main

import (
	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/internal/modules/auth"
	"github.com/LucasToledoPereira/go-boilerplate/internal/modules/collection"
	"github.com/LucasToledoPereira/go-boilerplate/internal/modules/game"
	"github.com/LucasToledoPereira/go-boilerplate/internal/modules/migration"
	"github.com/LucasToledoPereira/go-boilerplate/internal/modules/studio"
	"github.com/LucasToledoPereira/go-boilerplate/internal/modules/template"
	"github.com/LucasToledoPereira/go-boilerplate/internal/modules/user"
)

// @title Lakea API
// @version 1.0
// @description
// @host localhost:8080
// @BasePath /api/v1
func main() {
	builder := gbp.Default()
	builder.Register(migration.MigrationModuleFunc(), auth.AuthModuleFunc())
	builder.Register(studio.StudioModuleFunc())
	builder.Register(user.UserModuleFunc())
	builder.Register(game.GameModuleFunc())
	builder.Register(template.TemplateModuleFunc())
	builder.Register(collection.CollectionModuleFunc())
	server, _ := builder.Bootstrap()
	server.Run()
}
