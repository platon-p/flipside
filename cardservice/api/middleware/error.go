package middleware

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type errorMapper = func(error) int

var ErrBadRequest = errors.New("Bad request")
var ErrUnauthorized = errors.New("Unauthorized")

type ErrorMiddleware struct {
	mapper errorMapper
}

func NewErrorMiddleware(mappers ...errorMapper) *ErrorMiddleware {
	return &ErrorMiddleware{
		mapper: mappersChain(mappers...),
	}
}

func (mw *ErrorMiddleware) Handler(c *gin.Context) {
	c.Next()
	err := c.Errors.Last()
	if err == nil {
		return
	}
	log.Println(c.Request.URL.String(), err.Error())
	code := mw.mapper(err.Err)
	if code == -1 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
	} else {
		c.AbortWithStatusJSON(code, gin.H{
			"error": err.Error(),
		})
	}
}

func mappersChain(mapper ...errorMapper) errorMapper {
	return func(err error) int {
		for _, m := range mapper {
			if res := m(err); res != -1 {
				return res
			}
		}
		return -1
	}
}

func BasicErrorMapper(err error) int {
	switch {
	case errors.Is(err, strconv.ErrSyntax):
		return http.StatusBadRequest
	default:
		return -1
	}
}
