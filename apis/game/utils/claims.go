package utils

import (
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/config"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

func GetUserStudioID(ctx *gin.Context) (id uuid.UUID, err error) {
	claims := ctx.Value("JWT_PAYLOAD").(jwt.MapClaims)
	a := claims["studio"].(map[string]interface{})
	id, err = uuid.Parse(a["id"].(string))
	return
}

func GetUserIdentity(ctx *gin.Context) (identity string) {
	claims := ctx.Value("JWT_PAYLOAD").(jwt.MapClaims)
	identity = claims[config.C.Authorization.IdentityKey].(string)
	return
}

func GetUserRole(ctx *gin.Context) (role string) {
	claims := ctx.Value("JWT_PAYLOAD").(jwt.MapClaims)
	role = claims["type"].(string)
	return
}

func GetUserID(ctx *gin.Context) (id uuid.UUID, err error) {
	claims := ctx.Value("JWT_PAYLOAD").(jwt.MapClaims)
	id, err = uuid.Parse(claims["id"].(string))
	return id, err
}

func IsOwner(ctx *gin.Context) bool {
	role := GetUserRole(ctx)
	owner, _ := datatypes.OWNER.Value()
	return role == owner
}

func IsAdministrator(ctx *gin.Context) bool {
	role := GetUserRole(ctx)
	admin, _ := datatypes.ADMINISTRATOR.Value()
	return role == admin
}

func IsDeveloper(ctx *gin.Context) bool {
	role := GetUserRole(ctx)
	dev, _ := datatypes.DEVELOPER.Value()
	return role == dev
}

func IsLevelDesigner(ctx *gin.Context) bool {
	role := GetUserRole(ctx)
	designer, _ := datatypes.LEVEL_DESIGNER.Value()
	return role == designer
}

func IsCreator(ctx *gin.Context) bool {
	role := GetUserRole(ctx)
	creator, _ := datatypes.CREATOR.Value()
	return role == creator
}

func IsCommon(ctx *gin.Context) bool {
	role := GetUserRole(ctx)
	common, _ := datatypes.COMMON.Value()
	return role == common
}

func IsNotCommon(ctx *gin.Context) bool {
	return !IsCommon(ctx)
}

func IsOwnerOrAdministrator(ctx *gin.Context) bool {
	return IsOwner(ctx) || IsAdministrator(ctx)
}
