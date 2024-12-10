package route

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/controller"
	"github.com/platon-p/flipside/cardservice/api/helper"
	"github.com/platon-p/flipside/cardservice/api/middleware"
	"github.com/platon-p/flipside/cardservice/api/transfer"
	"github.com/platon-p/flipside/cardservice/repository"
	"github.com/platon-p/flipside/cardservice/service"
)

type CardRouter struct {
	controller      *controller.CardController
	authMiddleware  *middleware.AuthMiddleware
	errorMiddleware *middleware.ErrorMiddleware
}

func NewCardRouter(controller *controller.CardController, authMiddleware *middleware.AuthMiddleware) *CardRouter {
	return &CardRouter{
		controller:     controller,
		authMiddleware: authMiddleware,
	}
}

func (r *CardRouter) Setup(group *gin.RouterGroup) {
	cards := group.Group("/cards/:slug")
	cards.Use(r.errorMiddleware.Handler)
	cards.GET("/", r.GetCards)

	cards.Group("/").
		Use(r.authMiddleware.Handler()).
		POST("/", r.CreateCards).
		PUT("/", r.UpdateCards).
		DELETE("/", r.DeleteCards) // ?positions=...
}

func (r *CardRouter) CreateCards(ctx *gin.Context) {
	slug := ctx.Param("slug")
	var request []transfer.CardRequest
	if err := ctx.BindJSON(&request); err != nil {
		helper.ErrorMessage(ctx, http.StatusBadRequest, helper.BadRequest)
		return
	}
	userId := ctx.GetInt("userId")
	response, err := r.controller.CreateCards(userId, slug, request)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (r *CardRouter) GetCards(ctx *gin.Context) {
	slug := ctx.Param("slug")
	response, err := r.controller.GetCards(slug)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (r *CardRouter) UpdateCards(ctx *gin.Context) {
	slug := ctx.Param("slug")
	var request []transfer.CardRequest
	if err := ctx.BindJSON(&request); err != nil {
		helper.ErrorMessage(ctx, http.StatusBadRequest, helper.BadRequest)
		return
	}
	userId := ctx.GetInt("userId")
	response, err := r.controller.UpdateCards(userId, slug, request)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (r *CardRouter) DeleteCards(ctx *gin.Context) {
	slug := ctx.Param("slug")
	positions := strings.Split(ctx.Query("positions"), ",")
	userId := ctx.GetInt("userId")
	err := r.controller.DeleteCards(userId, slug, positions)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func cardErrorMapper(err error) int {
	switch {
	case errors.Is(err, repository.ErrCardSetNotFound) ||
		errors.Is(err, repository.ErrCardNotFound) ||

		errors.Is(err, repository.ErrCardWithThisPositionExists) ||
		errors.Is(err, service.ErrCardNegativePosition):
		return http.StatusBadRequest
	case errors.Is(err, service.ErrNotCardSetOwner):
		return http.StatusForbidden
	default:
		return -1
	}
}
