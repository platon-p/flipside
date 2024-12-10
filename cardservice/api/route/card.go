package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/helper"
	"github.com/platon-p/flipside/cardservice/api/middleware"
	"github.com/platon-p/flipside/cardservice/api/transfer"
	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/service"
)

type CardRouter struct {
	cardService    *service.CardService
	authMiddleware *middleware.AuthMiddleware
}

func NewCardRouter(cardService *service.CardService, authMiddleware *middleware.AuthMiddleware) *CardRouter {
	return &CardRouter{
		cardService:    cardService,
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

func (r *CardRouter) GetCards(ctx *gin.Context) {
	slug := ctx.Param("slug")
	models, err := r.cardService.GetCards(slug)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make([]transfer.CardResponse, len(models))
	for i := range models {
		resp[i] = cardModelToResponse(&models[i])
	}
	ctx.JSON(http.StatusOK, resp)
}

func (r *CardRouter) CreateCards(ctx *gin.Context) {
	slug := ctx.Param("slug")
	var request []transfer.CardRequest
	if err := ctx.BindJSON(&request); err != nil {
		helper.ErrorMessage(ctx, http.StatusBadRequest, helper.BadRequest)
		return
	}
	userId := ctx.GetInt("userId")
	models := make([]model.Card, len(request))
	for i := range request {
		models[i] = cardRequestToModel(&request[i])
	}
	modelsRes, err := r.cardService.CreateCards(userId, slug, models)
	if err != nil {
		ctx.Error(err)
		return
	}
	response := make([]transfer.CardResponse, len(modelsRes))
	for i := range modelsRes {
		response[i] = cardModelToResponse(&modelsRes[i])
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
	modelsReq := make([]model.Card, len(request))
	for i := range request {
		modelsReq[i] = cardRequestToModel(&request[i])
	}
	modelsResp, err := r.cardService.UpdateCards(userId, slug, modelsReq)
	if err != nil {
		ctx.Error(err)
		return
	}
	response := make([]transfer.CardResponse, len(modelsResp))
	for i := range modelsResp {
		response[i] = cardModelToResponse(&modelsResp[i])
	}
	ctx.JSON(http.StatusOK, response)
}

func (r *CardRouter) DeleteCards(ctx *gin.Context) {
	var req transfer.DeleteCardsRequest
	if err := ctx.ShouldBindQuery(req.Query); err != nil {
		ctx.Error(err)
		return
	}
	if err := ctx.ShouldBindUri(req.Uri); err != nil {
		ctx.Error(err)
		return
	}
	userId := ctx.GetInt("userId")
	err := r.cardService.DeleteCards(userId, req.Uri.Slug, req.Query.Positions)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func cardModelToResponse(card *model.Card) transfer.CardResponse {
	return transfer.CardResponse{
		Question:  card.Question,
		Answer:    card.Answer,
		Position:  card.Position,
		CardSetId: card.CardSetId,
	}
}

func cardRequestToModel(request *transfer.CardRequest) model.Card {
	return model.Card{
		Question: request.Question,
		Answer:   request.Answer,
		Position: request.Position,
	}
}
