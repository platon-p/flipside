package api

import (
	"net/http"
	"runtime/trace"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flashside/authservice/api/controller"
	"github.com/platon-p/flashside/authservice/api/transfer"
)

type CheckRouter struct {
	checkController *controller.CheckController
}

func (r *CheckRouter) Setup(group *gin.RouterGroup) {
	check := group.Group("/check")
	check.GET("/email/:email", r.CheckEmail)
	check.GET("/nickname/:nickname", r.CheckNickname)
}

func (r *CheckRouter) CheckEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	err := r.checkController.CheckEmail(email)
    if err != nil {
		ctx.JSON(http.StatusOK, transfer.MessageResponse{
            StatusCode: http.StatusOK,
            Message: "Ok",
        })
	} else {
        ctx.JSON(http.StatusBadRequest, transfer.MessageResponse{
            StatusCode: http.StatusBadRequest,
            Message: err.Error(),
        })
    }
}

func (r *CheckRouter) CheckNickname(ctx *gin.Context) {
	nickname := ctx.Param("nickname")
	err := r.checkController.CheckNickname(nickname)
    if err != nil {
		ctx.JSON(http.StatusOK, transfer.MessageResponse{
            StatusCode: http.StatusOK,
            Message: "Ok",
        })
	} else {
        ctx.JSON(http.StatusBadRequest, transfer.MessageResponse{
            StatusCode: http.StatusBadRequest,
            Message: err.Error(),
        })
    }
}
