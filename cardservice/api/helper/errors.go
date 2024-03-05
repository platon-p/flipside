package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/transfer"
)

var (
    BadRequest = "Bad request"
    Unauthorized = "Unauthorized"
)

func ErrorMessage(ctx *gin.Context, statusCode int, message string) {
    ctx.JSON(statusCode, transfer.MessageResponse{
    	StatusCode: statusCode,
    	Message:    message,
    })
}

