package auth_middleware

import (
	"log"
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/config"
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/LucasToledoPereira/go-boilerplate/internal/enums/codes"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	auth_commands "github.com/LucasToledoPereira/go-boilerplate/internal/modules/auth/commands"
	controller "github.com/LucasToledoPereira/go-boilerplate/internal/modules/auth/controller"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	middleware  *jwt.GinJWTMiddleware
	identityKey string
	controller  *controller.AuthController
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

/*
The above code initializes the authentication system for the application.
It creates a new instance of the Auth struct, sets the identity key and creates a middleware using the jwt package.
The middleware is configured with various options such as the realm, secret key, timeout, payload function, identity handler, authenticator, unauthorized handler, token lookup, token head name, time function, and whether to send a cookie.
If there is an error during initialization, the function logs a fatal error.
Finally, the function returns the Auth instance.
*/
func InitAuth() (a *Auth) {
	var err error
	a = &Auth{}
	a.identityKey = config.C.Authorization.IdentityKey
	a.middleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:           config.C.Authorization.Realm,
		Key:             []byte(config.C.Authorization.Secret),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     a.identityKey,
		PayloadFunc:     a.PayloadFunc,
		IdentityHandler: a.IdentityHandler,
		Authenticator:   a.Authenticator,
		Unauthorized:    a.Unauthorized,
		TokenLookup:     "header: Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		SendCookie:      true,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return a
}

func (a *Auth) SetUserController(uc *controller.AuthController) {
	a.controller = uc
}

func (a *Auth) PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(auth_commands.LoginCommandResponse); ok {
		return jwt.MapClaims{
			a.identityKey: v.Email,
			"id":          v.ID,
			"name":        v.Name,
			"nickname":    v.Nickname,
			"type":        v.Type,
			"image":       v.Image,
			"created_at":  v.CreatedAt,
			"studio": map[string]interface{}{
				"id":          v.Studio.ID,
				"name":        v.Studio.Name,
				"description": v.Studio.Description,
			},
		}
	}
	return jwt.MapClaims{}
}

func (a *Auth) IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &entity.User{
		Email: claims[a.identityKey].(string),
	}
}

func (a *Auth) Authenticator(c *gin.Context) (interface{}, error) {
	return a.controller.Login(c)
}

func (a *Auth) Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, models.NewResultWrapper(codes.NotAuthorized, false, nil, message))
}

// @tags Authentication
// Login User
// @Summary 	Return a JWT
// @Schemes
// @Description Return a JWT
// @Accept      json
// @Param		request	body	auth_commands.LoginCommandRequest	true	" "
// @Produce 	json
// @Success 	200
// @Router 		/login [post]
func (a *Auth) LoginHandler(c *gin.Context) {
	a.middleware.LoginHandler(c)
}

// @tags Authentication
// Logout User
// @Summary 	Logout
// @Schemes
// @Description Logout
// @Produce 	json
// @Success 	200
// @Router 		/logout [post]
func (a *Auth) LogoutHandler(c *gin.Context) {
	a.middleware.LogoutHandler(c)
}

// @tags Authentication
// Refresh Token
// @Summary 	Refresh JWT Token
// @Schemes
// @Description Refresh JWT Token
// @Produce 	json
// @Success 	200
// @Router 		/refresh [get]
func (a *Auth) RefreshHandler(c *gin.Context) {
	a.middleware.RefreshHandler(c)
}

func (a *Auth) MiddlewareFunc() gin.HandlerFunc {
	return a.middleware.MiddlewareFunc()
}
