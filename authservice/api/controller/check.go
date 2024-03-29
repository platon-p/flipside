package controller

import "github.com/platon-p/flipside/authservice/service"

type CheckController struct {
	checkService *service.CheckService
}

func NewCheckController(checkService *service.CheckService) *CheckController {
	return &CheckController{
		checkService: checkService,
	}
}

func (c *CheckController) CheckEmail(email string) error {
	return c.checkService.CheckEmail(email)
}

func (c *CheckController) CheckNickname(nickname string) error {
	return c.checkService.CheckNickname(nickname)
}
