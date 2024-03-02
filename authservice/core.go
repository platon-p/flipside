package main

import (
	"github.com/gin-gonic/gin"
	"github.com/platon-p/flashside/authservice/api/route"
)

type Core struct {
    Router *api.Router
}

func NewCore() *Core {
    router := api.NewRouter(&api.AuthRouter{}, &api.CheckRouter{})
    return &Core{
        Router: router,
    }
}

func (c *Core) Start() {
    r := gin.Default()
    c.Router.Setup(&r.RouterGroup)
    r.Run()
}
