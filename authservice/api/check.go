package api

import "github.com/gin-gonic/gin"

type CheckRouter struct{}

func (r *CheckRouter) Setup(group *gin.RouterGroup) {
	check := group.Group("/check")
	check.GET("/email/:email", r.CheckEmail)
	check.GET("/nickname/:nickname", r.CheckNickname)
}

func (r *CheckRouter) CheckEmail(ctx *gin.Context) {

}

func (r *CheckRouter) CheckNickname(ctx *gin.Context) {

}
