package route

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/controller"
	"github.com/platon-p/flipside/cardservice/api/helper"
	"github.com/platon-p/flipside/cardservice/service"
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
    switch {
    case errors.Is(err, service.ErrProfileNotFound):
        helper.ErrorMessage(ctx, http.StatusNotFound, err.Error())
    case err != nil:
        fmt.Println("GetProfile:", err)
        helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
    default:
        ctx.JSON(http.StatusOK, res)
    }
}

func (r *UserRouter) GetSetsHandler(ctx *gin.Context) {
    nickname := ctx.Param("nickname")
    res, err := r.controller.GetSets(nickname)
    switch {
    case errors.Is(err, service.ErrProfileNotFound):
        helper.ErrorMessage(ctx, http.StatusNotFound, err.Error())
    case err != nil:
        fmt.Println("GetSets:", err)
        helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
    default:
        ctx.JSON(http.StatusOK, res)
    }
}
