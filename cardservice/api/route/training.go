package route

import (
	"github.com/gin-gonic/gin"
	"github.com/platon-p/flipside/cardservice/api/middleware"
)

type TrainingRouter struct {
	authMiddleware *middleware.AuthMiddleware
}

func (r *TrainingRouter) Setup(group *gin.RouterGroup) {
	group.Use(r.authMiddleware.Handler()).
		GET("/cardset/:slug/trainings", r.GetCardSetTrainings).
		POST("/cardset/:slug/trainings", r.CreateTraining).
		GET("/training/:id/", r.GetTrainingSummary).
		GET("/training/:id/next", r.GetNextTask).
		POST("/training/:id/submit", r.SubmitTask)
}

func (r *TrainingRouter) GetCardSetTrainings(ctx *gin.Context)

func (r *TrainingRouter) CreateTraining(ctx *gin.Context)

func (r *TrainingRouter) GetTrainingSummary(ctx *gin.Context)

func (r *TrainingRouter) GetNextTask(ctx *gin.Context)

func (r *TrainingRouter) SubmitTask(ctx *gin.Context)
