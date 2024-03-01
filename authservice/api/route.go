package api

import "github.com/gin-gonic/gin"

func AddRoutes(group *gin.RouterGroup) {
    api := group.Group("/api")

    auth := api.Group("/auth")
    auth.POST("/register", RegisterHandler)
    auth.POST("/login-by-email", LoginByEmailHandler)
    auth.POST("/login-by-token", LoginByTokenHandler)
}

func RegisterHandler(ctx *gin.Context) {

}

func LoginByEmailHandler(ctx *gin.Context) {

}

func LoginByTokenHandler(ctx *gin.Context) {

}
