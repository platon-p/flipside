package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flashside/authservice/api/controller"
	"github.com/platon-p/flashside/authservice/api/transfer"
	"github.com/platon-p/flashside/authservice/service"
)

type AuthRouter struct{
    controller *controller.AuthController
}

func (r *AuthRouter) Setup(group *gin.RouterGroup) {
	auth := group.Group("/auth")
	auth.POST("/register", r.Register)
	auth.POST("/login-by-email", r.LoginByEmail)
	auth.POST("/login-by-token", r.LoginByToken)
}

func (r *AuthRouter) Register(ctx *gin.Context) {
    var request transfer.RegisterRequest
    if err := ctx.BindJSON(&request); err != nil {
        return
    }
    res, err := r.controller.Register(request)
    switch err {
    case nil:
        ctx.JSON(http.StatusOK, res)
    case service.NicknameExistsError, service.EmailExistsError, service.EmailIncorrectFormatError, service.NicknameIncorrectFormatError:
        ctx.JSON(http.StatusBadRequest, transfer.ErrorResponse{
        	StatusCode: http.StatusBadRequest,
        	Message:    err.Error(),
        })
    default:
        fmt.Println("Internal error ", err)
        ctx.JSON(http.StatusInternalServerError, transfer.ErrorResponse{
            StatusCode: http.StatusInternalServerError,
            Message: "Internal server error",
        })
    }
}

func (r *AuthRouter) LoginByEmail(ctx *gin.Context) {
    var request transfer.LoginByEmailRequest
    if err := ctx.BindJSON(&request); err != nil {
        return
    }
    res, err := r.controller.LoginByEmail(request)
    switch err {
    case nil:
        ctx.JSON(http.StatusOK, res)
    case service.BadCredentialsError:
        ctx.JSON(http.StatusBadRequest, transfer.ErrorResponse{
        	StatusCode: http.StatusBadRequest,
        	Message:    err.Error(),
        })
    default:
        fmt.Println("Internal error ", err)
        ctx.JSON(http.StatusInternalServerError, transfer.ErrorResponse{
            StatusCode: http.StatusInternalServerError,
            Message: "Internal server error",
        })
    }
}

func (r *AuthRouter) LoginByToken(ctx *gin.Context) {
    var request transfer.LoginByTokenRequest
    if err := ctx.BindJSON(&request); err != nil {
        return
    }
    res, err := r.controller.LoginByToken(request)
    switch err {
    case nil:
        ctx.JSON(http.StatusOK, res)
    case service.InvalidRefreshToken, service.ExpiredRefreshToken:
        ctx.JSON(http.StatusBadRequest, transfer.ErrorResponse{
        	StatusCode: http.StatusBadRequest,
        	Message:    err.Error(),
        })
    default:
        fmt.Println("Internal error ", err)
        ctx.JSON(http.StatusInternalServerError, transfer.ErrorResponse{
            StatusCode: http.StatusInternalServerError,
            Message: "Internal server error",
        })
    }
}
