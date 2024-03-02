package controller

import "github.com/platon-p/flashside/authservice/service"

type CheckController struct {
    checkService *service.CheckService
}

func (c *CheckController) CheckEmail(email string) error {
    return c.checkService.CheckEmail(email)
}

func (c *CheckController) CheckNickname(nickname string) error {
    return c.checkService.CheckNickname(nickname)
}
