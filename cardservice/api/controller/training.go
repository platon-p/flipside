package controller

import (
	"github.com/platon-p/flipside/cardservice/api/transfer"
	"github.com/platon-p/flipside/cardservice/service/training"
)

type TrainingController struct {
	service *training.TrainingService
}

func (c *TrainingController) GetCardSetTrainings(userId int, slug string) ([]transfer.TrainingSummaryResponse, error)

func (c *TrainingController) CreateTraining(userId int, slug string, trainingType string) (*transfer.TrainingSummaryResponse, error)

func (c *TrainingController) GetTrainingSummary(userId int, trainingId string) (*transfer.TrainingSummaryResponse, error)

func (c *TrainingController) GetNextTask(userId int, trainingId string) (*transfer.TaskResponse, error)

func (c *TrainingController) SubmitTask(userId int, trainingId string, answer string) (*transfer.TaskResultResponse, error)
