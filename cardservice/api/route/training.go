package route

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/controller"
	"github.com/platon-p/flipside/cardservice/api/helper"
	"github.com/platon-p/flipside/cardservice/api/middleware"
	"github.com/platon-p/flipside/cardservice/service"
	"github.com/platon-p/flipside/cardservice/service/training"
)

type TrainingRouter struct {
	controller     *controller.TrainingController
	authMiddleware *middleware.AuthMiddleware
}

func NewTrainingRouter(controller *controller.TrainingController, authMiddleware *middleware.AuthMiddleware) *TrainingRouter {
	return &TrainingRouter{
		controller:     controller,
		authMiddleware: authMiddleware,
	}
}

func (r *TrainingRouter) Setup(group *gin.RouterGroup) {
	group.Use(r.authMiddleware.Handler()).
		GET("/cardset/:slug/trainings", r.GetCardSetTrainings).
		POST("/cardset/:slug/trainings", r.CreateTraining). // ?type=string
		GET("/training/:id/", r.GetTrainingSummary).
		GET("/training/:id/next", r.GetNextTask).
		POST("/training/:id/submit", r.SubmitTask) // ?answer=string
}

func (r *TrainingRouter) GetCardSetTrainings(ctx *gin.Context) {
	slug := ctx.Param("slug")
	userId := ctx.GetInt("userId")
	res, err := r.controller.GetCardSetTrainings(userId, slug)
	switch {
    case errors.Is(err, service.ErrCardSetNotFound):
        helper.ErrorMessage(ctx, http.StatusNotFound, err.Error())
    case errors.Is(err, training.ErrNotATrainingOwner):
        helper.ErrorMessage(ctx, http.StatusForbidden, err.Error())
	case err != nil:
		fmt.Println("GetCardSetTrainings:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, helper.InternalServerError)
	default:
		ctx.JSON(http.StatusOK, res)
	}
}

func (r *TrainingRouter) CreateTraining(ctx *gin.Context) {
	slug := ctx.Param("slug")
	userId := ctx.GetInt("userId")
	trainingType := ctx.Query("type")
	res, err := r.controller.CreateTraining(userId, slug, trainingType)
	switch {
    case errors.Is(err, controller.ErrUnresolvedTrainingType):
        helper.ErrorMessage(ctx, http.StatusBadRequest, err.Error())
    case errors.Is(err, service.ErrCardSetNotFound):
        helper.ErrorMessage(ctx, http.StatusNotFound, err.Error())
	case err != nil:
		fmt.Println("CreateTraining:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, helper.InternalServerError)
	default:
		ctx.JSON(http.StatusOK, res)
	}
}

func (r *TrainingRouter) GetTrainingSummary(ctx *gin.Context) {
	trainingId := ctx.Param("id")
	userId := ctx.GetInt("userId")
	res, err := r.controller.GetTrainingSummary(userId, trainingId)
	switch {
	case errors.Is(err, strconv.ErrSyntax):
		helper.ErrorMessage(ctx, http.StatusBadRequest, helper.BadRequest)
    case errors.Is(err, training.ErrNotATrainingOwner):
        helper.ErrorMessage(ctx, http.StatusForbidden, err.Error())
    case errors.Is(err, training.ErrTrainingNotFound):
        helper.ErrorMessage(ctx, http.StatusNotFound, err.Error())
	case err != nil:
		fmt.Println("GetTrainingSummary:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, helper.InternalServerError)
	default:
		ctx.JSON(http.StatusOK, res)
	}
}

func (r *TrainingRouter) GetNextTask(ctx *gin.Context) {
	trainingId := ctx.Param("id")
	userId := ctx.GetInt("userId")
	res, err := r.controller.GetNextTask(userId, trainingId)
	switch {
	case errors.Is(err, strconv.ErrSyntax):
		helper.ErrorMessage(ctx, http.StatusBadRequest, helper.BadRequest)
	case err != nil:
		fmt.Println("GetNextTask:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, helper.InternalServerError)
	default:
		ctx.JSON(http.StatusOK, res)
	}
}

func (r *TrainingRouter) SubmitTask(ctx *gin.Context) {
	trainingId := ctx.Param("id")
	userId := ctx.GetInt("userId")
	answer, answerFound := ctx.GetQuery("answer")
	if !answerFound {
		helper.ErrorMessage(ctx, http.StatusBadRequest, helper.BadRequest)
		return
	}
	res, err := r.controller.SubmitTask(userId, trainingId, answer)
	switch {
	case errors.Is(err, strconv.ErrSyntax):
		helper.ErrorMessage(ctx, http.StatusBadRequest, helper.BadRequest)
	case err != nil:
		fmt.Println("SubmitTask:", err)
		helper.ErrorMessage(ctx, http.StatusInternalServerError, helper.InternalServerError)
	default:
		ctx.JSON(http.StatusOK, res)
	}
}
