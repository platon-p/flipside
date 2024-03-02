package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flashside/authservice/api/controller"
	"github.com/platon-p/flashside/authservice/api/transfer"
	"github.com/platon-p/flashside/authservice/service"
)

type CheckRouter struct {
	checkController *controller.CheckController
}

func NewCheckRouter(checkController *controller.CheckController) *CheckRouter {
    return &CheckRouter{
        checkController: checkController,
    }
}

func (r *CheckRouter) Setup(group *gin.RouterGroup) {
	check := group.Group("/check")
	check.GET("/email/:email", r.CheckEmail)
	check.GET("/nickname/:nickname", r.CheckNickname)
}

func (r *CheckRouter) CheckEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	err := r.checkController.CheckEmail(email)
    switch err {
    case nil:
		ctx.JSON(http.StatusOK, transfer.MessageResponse{
            StatusCode: http.StatusOK,
            Message: "Ok",
        })
    case service.EmailExistsError:
        ctx.JSON(http.StatusBadRequest, transfer.MessageResponse{
            StatusCode: http.StatusBadRequest,
            Message: err.Error(),
        })
    default:
        fmt.Println("Internal error ", err)
        ctx.JSON(http.StatusInternalServerError, transfer.MessageResponse{
            StatusCode: http.StatusInternalServerError,
            Message: "Internal server error",
        })
    }
}

func (r *CheckRouter) CheckNickname(ctx *gin.Context) {
	nickname := ctx.Param("nickname")
	err := r.checkController.CheckNickname(nickname)
    switch err {
    case nil:
		ctx.JSON(http.StatusOK, transfer.MessageResponse{
            StatusCode: http.StatusOK,
            Message: "Ok",
        })
    case service.NicknameExistsError:
        ctx.JSON(http.StatusBadRequest, transfer.MessageResponse{
            StatusCode: http.StatusBadRequest,
            Message: err.Error(),
        })
    default:
        fmt.Println("Internal error ", err)
        ctx.JSON(http.StatusInternalServerError, transfer.MessageResponse{
            StatusCode: http.StatusInternalServerError,
            Message: "Internal server error",
        })
    }
}
