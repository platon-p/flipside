package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/transfer"
	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/service"
)

type UserRouter struct {
	userService *service.UserService
}

func NewUserRouter(userService *service.UserService) *UserRouter {
	return &UserRouter{
		userService: userService,
	}
}

func (r *UserRouter) Setup(group *gin.RouterGroup) {
	group.Group("/users").
		GET("/:nickname/profile", r.GetProfileHandler).
		GET("/:nickname/sets", r.GetSetsHandler)
}

func (r *UserRouter) GetProfileHandler(ctx *gin.Context) {
	nickname := ctx.Param("nickname")
	res, err := r.userService.GetProfile(nickname)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := profileResponseFromModel(res)
	ctx.JSON(http.StatusOK, resp)
}

func (r *UserRouter) GetSetsHandler(ctx *gin.Context) {
	nickname := ctx.Param("nickname")
	models, err := r.userService.GetCardSets(nickname)
	if err != nil {
		ctx.Error(err)
		return
	}
	res := make([]transfer.CardSetResponse, len(models))
	for i := range models {
		res[i] = cardSetModelToResponse(&models[i])
	}
	ctx.JSON(http.StatusOK, res)
}

func profileResponseFromModel(model *model.Profile) transfer.ProfileResponse {
	return transfer.ProfileResponse{
		Id:       model.Id,
		Name:     model.Name,
		Nickname: model.Nickname,
	}
}

func cardSetModelToResponse(cardSet *model.CardSet) transfer.CardSetResponse {
	return transfer.CardSetResponse{
		Title:   cardSet.Title,
		Slug:    cardSet.Slug,
		OwnerId: cardSet.OwnerId,
	}
}
