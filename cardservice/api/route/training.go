package route

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/middleware"
	"github.com/platon-p/flipside/cardservice/api/transfer"
	"github.com/platon-p/flipside/cardservice/repository"
	"github.com/platon-p/flipside/cardservice/service/training"
)

type TrainingRouter struct {
	service        *training.TrainingService
	authMiddleware *middleware.AuthMiddleware
}

func NewTrainingRouter(service *training.TrainingService, authMiddleware *middleware.AuthMiddleware) *TrainingRouter {
	return &TrainingRouter{
		service:        service,
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
	modelsResp, err := r.service.GetCardSetTrainings(userId, slug)
	if err != nil {
		ctx.Error(err)
		return
	}
	res := transfer.TrainingSummariesToResponse(modelsResp)
	ctx.JSON(http.StatusOK, res)
}

func (r *TrainingRouter) CreateTraining(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	slug := ctx.Param("slug")
	trainingTypeStr := ctx.Query("type")
	trainingType, err := transfer.ResolveTrainingType(trainingTypeStr)
	if err != nil {
		ctx.Error(err)
		return
	}
	model, err := r.service.CreateTraining(userId, slug, trainingType)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := transfer.TrainingSummaryToResponse(*model)
	ctx.JSON(http.StatusOK, resp)
}

func (r *TrainingRouter) GetTrainingSummary(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	trainingId := ctx.Param("id")
	trainingIdInt, err := strconv.Atoi(trainingId)
	if err != nil {
		ctx.Error(err)
		return
	}
	model, err := r.service.GetTrainingSummary(userId, trainingIdInt)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := transfer.TrainingSummaryToResponse(*model)
	ctx.JSON(http.StatusOK, resp)
}

func (r *TrainingRouter) GetNextTask(ctx *gin.Context) {
	trainingId := ctx.Param("id")
	userId := ctx.GetInt("userId")
	trainingIdInt, err := strconv.Atoi(trainingId)
	if err != nil {
		ctx.Error(err)
		return
	}
	taskModel, err := r.service.GetNextTask(userId, trainingIdInt)
	if err != nil {
		ctx.Error(err)
		return
	}
	res := transfer.TaskToResponse(taskModel)
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
	trainingIdInt, err := strconv.Atoi(trainingId)
	if err != nil {
		ctx.Error(err)
		return
	}
	taskResultModel, err := r.service.SubmitAnswer(userId, trainingIdInt, answer)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, transfer.TaskResultResponse{
		Answer:    *taskResultModel.Answer,
		IsCorrect: *taskResultModel.IsCorrect,
	})
}

func trainingsErrorMapper(err error) int {
	switch {
	case errors.Is(err, training.ErrTrainingNotFound) ||
		errors.Is(err, repository.ErrCardSetNotFound):
		return http.StatusNotFound
	case errors.Is(err, training.ErrTrainingIsCompleted) ||
		errors.Is(err, training.ErrInvalidAnswer) ||
		errors.Is(err, training.ErrTaskNotFound) ||
		errors.Is(err, transfer.ErrUnresolvedTrainingType):
		return http.StatusBadRequest
	case errors.Is(err, training.ErrNotATrainingOwner):
		return http.StatusForbidden
	default:
		return -1
	}
}
