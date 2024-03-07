package route

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/controller"
	"github.com/platon-p/flipside/cardservice/api/helper"
	"github.com/platon-p/flipside/cardservice/api/middleware"
	"github.com/platon-p/flipside/cardservice/api/transfer"
	"github.com/platon-p/flipside/cardservice/repository"
	"github.com/platon-p/flipside/cardservice/service"
)

type CardSetRouter struct {
	controller     *controller.CardSetController
	authMiddleware *middleware.AuthMiddleware
}

func NewCardSetRouter(controller *controller.CardSetController, authMiddleware *middleware.AuthMiddleware) *CardSetRouter {
	return &CardSetRouter{
		controller:     controller,
		authMiddleware: authMiddleware,
	}
}

func (r *CardSetRouter) Setup(group *gin.RouterGroup) {
	cardset := group.Group("/cardset")

	cardset.GET("/:slug", r.GetCardSet)
	cardset.Group("/").
		Use(r.authMiddleware.Handler()).
		POST("/", r.CreateCardSet).
		PUT("/:slug", r.UpdateCardSet).
		DELETE("/:slug", r.DeleteCardSet)
}

func (r *CardSetRouter) CreateCardSet(ctx *gin.Context) {
	var request transfer.CardSetRequest
	if err := ctx.BindJSON(&request); err != nil {
		helper.ErrorMessage(ctx, http.StatusBadRequest, helper.BadRequest)
		return
	}
	userId := ctx.GetInt("userId")
	response, err := r.controller.CreateCardSet(userId, &request)
	switch {
	case errors.Is(err, repository.ErrCardSetSlugAlreadyExists):
		helper.ErrorMessage(ctx, http.StatusBadRequest, err.Error())
	case err != nil:
		fmt.Println("CreateCardSet:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
	default:
		ctx.JSON(http.StatusCreated, response)
	}
}

func (r *CardSetRouter) GetCardSet(ctx *gin.Context) {
	slug := ctx.Param("slug")
	response, err := r.controller.GetCardSet(slug)
	if errors.Is(err, service.ErrCardSetNotFound) {
		helper.ErrorMessage(ctx, http.StatusNotFound, err.Error())
	} else if err != nil {
		fmt.Println("GetCardSet:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
	} else {
		ctx.JSON(http.StatusOK, response)
	}
}

func (r *CardSetRouter) UpdateCardSet(ctx *gin.Context) {
	slug := ctx.Param("slug")
	var request transfer.CardSetRequest
	if err := ctx.BindJSON(&request); err != nil {
		helper.ErrorMessage(ctx, http.StatusBadRequest, helper.BadRequest)
		return
	}
	userId := ctx.GetInt("userId")
	response, err := r.controller.UpdateCardSet(userId, slug, &request)
	switch {
	case errors.Is(err, service.ErrNotCardSetOwner):
		helper.ErrorMessage(ctx, http.StatusForbidden, err.Error())
	case errors.Is(err, service.ErrCardSetNotFound):
		helper.ErrorMessage(ctx, http.StatusNotFound, err.Error())
	case err != nil:
		fmt.Println("UpdateCardSet:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
	default:
		ctx.JSON(http.StatusOK, response)
	}
}

func (r *CardSetRouter) DeleteCardSet(ctx *gin.Context) {
	slug := ctx.Param("slug")
	userId := ctx.GetInt("userId")
	if err := r.controller.DeleteCardSet(userId, slug); errors.Is(err, service.ErrCardSetNotFound) {
		helper.ErrorMessage(ctx, http.StatusNotFound, err.Error())
	} else if err != nil {
		fmt.Println("DeleteCardSet:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
	} else {
		helper.ErrorMessage(ctx, http.StatusOK, "Success")
	}
}
