package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/controller"
)

type UserRouter struct {
	controller *controller.UserController
}

func NewUserRouter(userController *controller.UserController) *UserRouter {
	return &UserRouter{
		controller: userController,
	}
}

func (r *UserRouter) Setup(group *gin.RouterGroup) {
	group.Group("/users").
		GET("/:nickname/profile", r.GetProfileHandler).
		GET("/:nickname/sets", r.GetSetsHandler)
}

func (r *UserRouter) GetProfileHandler(ctx *gin.Context) {
	nickname := ctx.Param("nickname")
	res, err := r.controller.GetProfile(nickname)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (r *UserRouter) GetSetsHandler(ctx *gin.Context) {
	nickname := ctx.Param("nickname")
	res, err := r.controller.GetSets(nickname)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
