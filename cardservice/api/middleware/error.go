package middleware

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/repository"
	"github.com/platon-p/flipside/cardservice/service"
	"github.com/platon-p/flipside/cardservice/service/training"
)

var ErrBadRequest = errors.New("Bad request")

var chain = mappersChain(profileMap, trainingMap, cardMap, cardsetMap, otherMap)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}
		log.Println(c.Request.URL.String(), err.Error())
		code := chain(err.Err)
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
}

func mappersChain(mapper ...func(error) int) func(error) int {
	return func(err error) int {
		for _, m := range mapper {
			if res := m(err); res != -1 {
				return res
			}
		}
		return -1
	}
}

func profileMap(err error) int {
	switch {
	case errors.Is(err, service.ErrProfileNotFound):
		return http.StatusNotFound
	default:
		return -1
	}
}

func trainingMap(err error) int {
	switch {
	case errors.Is(err, training.ErrTrainingNotFound):
		return http.StatusNotFound
	case errors.Is(err, training.ErrTrainingIsCompleted) ||
		errors.Is(err, training.ErrInvalidAnswer) ||
		errors.Is(err, training.ErrTaskNotFound):
		return http.StatusBadRequest
	case errors.Is(err, training.ErrNotATrainingOwner):
		return http.StatusForbidden
	default:
		return -1
	}
}

func otherMap(err error) int {
	switch {
	case errors.Is(err, strconv.ErrSyntax):
		return http.StatusBadRequest
	default:
		return -1
	}
}

func cardsetMap(err error) int {
	switch {
	case errors.Is(err, repository.ErrCardSetSlugAlreadyExists):
		return http.StatusBadRequest
	case errors.Is(err, repository.ErrCardSetNotFound):
		return http.StatusNotFound
	case errors.Is(err, service.ErrNotCardSetOwner):
		return http.StatusForbidden
	default:
		return -1
	}
}

func cardMap(err error) int {
	switch {
	case errors.Is(err, repository.ErrCardSetNotFound) ||
		errors.Is(err, repository.ErrCardNotFound) ||

		errors.Is(err, repository.ErrCardWithThisPositionExists) ||
		errors.Is(err, service.ErrCardNegativePosition):
		return http.StatusBadRequest
	case errors.Is(err, service.ErrNotCardSetOwner):
		return http.StatusForbidden
	default:
		return -1
	}
}
