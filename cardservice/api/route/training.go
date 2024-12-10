package route

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/controller"
	"github.com/platon-p/flipside/cardservice/api/middleware"
	"github.com/platon-p/flipside/cardservice/repository"
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
	mw := middleware.NewErrorMiddleware(trainingsErrorMapper)
    trainings := group.Group("")
    trainings.
		Use(r.authMiddleware.Handler()).
		Use(mw.Handler).
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
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (r *TrainingRouter) CreateTraining(ctx *gin.Context) {
	slug := ctx.Param("slug")
	userId := ctx.GetInt("userId")
	trainingType := ctx.Query("type")
	res, err := r.controller.CreateTraining(userId, slug, trainingType)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (r *TrainingRouter) GetTrainingSummary(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	trainingId := ctx.Param("id")
	res, err := r.controller.GetTrainingSummary(userId, trainingId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (r *TrainingRouter) GetNextTask(ctx *gin.Context) {
	trainingId := ctx.Param("id")
	userId := ctx.GetInt("userId")
	res, err := r.controller.GetNextTask(userId, trainingId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (r *TrainingRouter) SubmitTask(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	trainingId := ctx.Param("id")
	answer, answerFound := ctx.GetQuery("answer")
	if !answerFound {
		ctx.Error(middleware.ErrBadRequest)
		return
	}
	res, err := r.controller.SubmitTask(userId, trainingId, answer)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func trainingsErrorMapper(err error) int {
	switch {
	case errors.Is(err, training.ErrTrainingNotFound) ||
		errors.Is(err, repository.ErrCardSetNotFound):
		return http.StatusNotFound
	case errors.Is(err, training.ErrTrainingIsCompleted) ||
		errors.Is(err, training.ErrInvalidAnswer) ||
		errors.Is(err, training.ErrTaskNotFound) ||
		errors.Is(err, controller.ErrUnresolvedTrainingType):
		return http.StatusBadRequest
	case errors.Is(err, training.ErrNotATrainingOwner):
		return http.StatusForbidden
	default:
		return -1
	}
}
