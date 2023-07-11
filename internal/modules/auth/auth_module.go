package auth

import (
	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/config"
	controller "github.com/LucasToledoPereira/go-boilerplate/internal/modules/auth/controller"
	handler "github.com/LucasToledoPereira/go-boilerplate/internal/modules/auth/handler"
	middleware "github.com/LucasToledoPereira/go-boilerplate/internal/modules/auth/middleware"
	repository "github.com/LucasToledoPereira/go-boilerplate/internal/modules/auth/repository"

	"github.com/LucasToledoPereira/go-boilerplate/internal/router"
)

func AuthModuleFunc() gbp.ModuleFunc {
	return func(input *gbp.ModuleFuncInput) (err error) {
		ar := repository.NewAuthRepository(input.Datastore)
		ah := handler.NewAuthHandler(&ar, input.Filestore)
		ac := controller.NewAuthController(ah)
		createRouters(input.Router, ac)
		return nil
	}
}

func createRouters(r *router.Router, c *controller.AuthController) {
	authm := middleware.InitAuth()
	r.Private.Use(authm.MiddlewareFunc())
	authm.SetUserController(c)
	pb := r.Public
	if config.C.Settings.MultiTenant {
		pb.POST("register", c.Register)
	}
	pb.POST("login", authm.LoginHandler)
	pb.POST("logout", authm.LogoutHandler)
	pb.POST("refresh", authm.RefreshHandler)
}
