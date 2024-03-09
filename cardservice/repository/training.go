package repository

import "github.com/platon-p/flipside/cardservice/model"

type TrainingRepository interface {
    CreateTraining(training *model.Training) (*model.Training, error)
    GetTraining(trainingId int) (*model.Training, error)
    SetTrainingStatus(trainingId int, status string) (*model.Training, error)
    GetTaskResults(trainingId int) ([]model.TrainingTaskResult, error)
    CreateTaskResult(trainingId int, taskResult *model.TrainingTaskResult) (*model.TrainingTaskResult, error)
}
