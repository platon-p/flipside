package repository

import "github.com/platon-p/flipside/cardservice/model"

type TrainingRepository interface {
	CreateTraining(training *model.Training) (*model.TrainingSummary, error)
	GetTraining(trainingId int) (*model.Training, error)
    GetCardSetTrainings(userId int, cardSetId int) ([]model.TrainingSummary, error)
	SetTrainingStatus(trainingId int, status string) (*model.TrainingSummary, error)
	GetTaskResults(trainingId int) ([]model.TrainingTaskResult, error)
    GetLastTaskResult(trainingId int) (*model.TrainingTaskResult, error)
	CreateTaskResult(trainingId int, taskResult *model.TrainingTaskResult) (*model.TrainingTaskResult, error)
}
