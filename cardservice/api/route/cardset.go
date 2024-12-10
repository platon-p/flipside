package route

import (
	"errors"
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
	mw := middleware.NewErrorMiddleware(cardErrorMapper)

	cardset := group.Group("/cardset")
	cardset.
		Use(mw.Handler).
		GET("/:slug", r.GetCardSet)
    // TODO: check if the grouping is correct
	cardset.Group("/").
		Use(r.authMiddleware.Handler()).
		POST("/", r.CreateCardSet).
		PUT("/:slug", r.UpdateCardSet).
		DELETE("/:slug", r.DeleteCardSet)
}

func (r *CardSetRouter) GetCardSet(ctx *gin.Context) {
	slug := ctx.Param("slug")
	response, err := r.controller.GetCardSet(slug)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (r *CardSetRouter) CreateCardSet(ctx *gin.Context) {
	var request transfer.CardSetRequest
	if err := ctx.BindJSON(&request); err != nil {
		helper.ErrorMessage(ctx, http.StatusBadRequest, helper.BadRequest)
		return
	}
	userId := ctx.GetInt("userId")
	response, err := r.controller.CreateCardSet(userId, &request)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response)
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
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (r *CardSetRouter) DeleteCardSet(ctx *gin.Context) {
	slug := ctx.Param("slug")
	userId := ctx.GetInt("userId")
	err := r.controller.DeleteCardSet(userId, slug)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func cardsetErrorMapper(err error) int {
	switch {
	case errors.Is(err, repository.ErrCardSetSlugAlreadyExists):
		return http.StatusBadRequest
	case errors.Is(err, repository.ErrCardSetNotFound):
		return http.StatusNotFound
	case errors.Is(err, service.ErrNotCardSetOwner):
		return http.StatusForbidden
	default:
		return -1
	}
}
