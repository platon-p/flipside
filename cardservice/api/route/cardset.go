package route

import (
	"fmt"
	"net/http"
	"strconv"

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

func (r *CardSetRouter) Setup(group *gin.RouterGroup) {
	cardset := group.Group("/cardset")

	cardset.GET("/:slug", r.GetCardSet)
	cardset.Group("/").
		Use(r.authMiddleware.Handler()).
		POST("/", r.CreateCardSet).
		PUT("/", r.UpdateCardSet).
		DELETE("/:id", r.DeleteCardSet)
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
		fmt.Println("CreateCardSet:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
	} else {
		ctx.JSON(http.StatusCreated, response)
	}
}

func (r *CardSetRouter) GetCardSet(ctx *gin.Context) {
	slug := ctx.Param("slug")
	response, err := r.controller.GetCardSet(slug)
	if err != nil {
		fmt.Println("GetCardSet:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
	} else {
		ctx.JSON(http.StatusOK, response)
	}
}

func (r *CardSetRouter) UpdateCardSet(ctx *gin.Context) {
	var request transfer.CardSetRequest
	if err := ctx.BindJSON(&request); err != nil {
		helper.ErrorMessage(ctx, http.StatusBadRequest, helper.BadRequest)
		return
	}
	userId := ctx.GetInt("userId")
	response, err := r.controller.UpdateCardSet(userId, &request)
	if err != nil {
		fmt.Println("UpdateCardSet:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
	} else {
		ctx.JSON(http.StatusOK, response)
	}

}
func (r *CardSetRouter) DeleteCardSet(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorMessage(ctx, http.StatusBadRequest, helper.BadRequest)
		return
	}
	userId := ctx.GetInt("userId")
	if err := r.controller.DeleteCardSet(userId, id); err != nil {
		fmt.Println("DeleteCardSet:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, "Internal server error")
	} else {
		helper.ErrorMessage(ctx, http.StatusOK, "Success")
	}
}
