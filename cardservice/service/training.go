package service

import "github.com/platon-p/flipside/cardservice/model"

type TrainingService struct {
}

func (s *TrainingService) CreateTraining(userId int, slug string, trainingType string) (*model.Training, error)

func (s *TrainingService) GetTrainingSummary(userId int, trainingId int) (*model.TrainingSummary, error)

func (s *TrainingService) GetNextTask(userId int, trainingId int) (*model.Task, error)

func (s *TrainingService) SubmitAnswer(userId int, trainingId int, answer string) error
