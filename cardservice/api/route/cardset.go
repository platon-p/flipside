package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/controller"
	"github.com/platon-p/flipside/cardservice/api/helper"
	"github.com/platon-p/flipside/cardservice/api/middleware"
	"github.com/platon-p/flipside/cardservice/api/transfer"
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
