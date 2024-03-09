package controller

import (
	"github.com/platon-p/flipside/cardservice/api/transfer"
	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) GetProfile(nickname string) (*transfer.ProfileResponse, error) {
	model, err := c.userService.GetProfile(nickname)
	if err != nil {
		return nil, err
	}
	return &transfer.ProfileResponse{
		Id:       model.Id,
		Name:     model.Name,
		Nickname: model.Nickname,
	}, nil
}

func (c *UserController) GetSets(nickname string) ([]transfer.CardSetResponse, error) {
	models, err := c.userService.GetCardSets(nickname)
	if err != nil {
		return nil, err
	}
	res := make([]transfer.CardSetResponse, len(models))
	for i := range models {
		res[i] = *c.cardSetModelToResponse(&models[i])
	}
	return res, nil
}

func (c *UserController) cardSetModelToResponse(cardSet *model.CardSet) *transfer.CardSetResponse {
	return &transfer.CardSetResponse{
		Title:   cardSet.Title,
		Slug:    cardSet.Title,
		OwnerId: cardSet.OwnerId,
	}
}
