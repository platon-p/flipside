package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/controller"
	"github.com/platon-p/flipside/cardservice/api/helper"
	"github.com/platon-p/flipside/cardservice/api/middleware"
	"github.com/platon-p/flipside/cardservice/api/transfer"
)

type CardRouter struct {
	controller *controller.CardController
    authMiddleware *middleware.AuthMiddleware
}

func NewCardRouter(controller *controller.CardController, authMiddleware *middleware.AuthMiddleware) *CardRouter {
	return &CardRouter{
		controller: controller,
        authMiddleware: authMiddleware,
	}
}

func (r *CardRouter) Setup(group *gin.RouterGroup) {
	cards := group.Group("/cards/:slug")
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
		fmt.Println("CreateCards:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
	} else {
		ctx.JSON(http.StatusOK, response)
	}
}
func (r *CardRouter) GetCards(ctx *gin.Context) {
    slug := ctx.Param("slug")
    response, err := r.controller.GetCards(slug)
	if err != nil {
		fmt.Println("GetCards:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
	} else {
		ctx.JSON(http.StatusOK, response)
	}
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
		fmt.Println("UpdateCards:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
	} else {
		ctx.JSON(http.StatusOK, response)
	}
}
func (r *CardRouter) DeleteCards(ctx *gin.Context) {
    slug := ctx.Param("slug")
    positions := ctx.QueryArray("positions")
    userId := ctx.GetInt("userId")
    err := r.controller.DeleteCards(userId, slug, positions)
	if err != nil {
		fmt.Println("DeleteCards:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
	} else {
		helper.ErrorMessage(ctx, http.StatusOK, "Success")
	}
}
