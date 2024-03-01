package main

import (
	"github.com/gin-gonic/gin"
	"github.com/platon-p/flashside/authservice/api"
)

func main() {
    r := gin.Default()
    api.AddRoutes(&r.RouterGroup)
    r.Run()
}
