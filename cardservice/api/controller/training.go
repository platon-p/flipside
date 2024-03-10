package controller

import (
	"errors"
	"strconv"

	"github.com/platon-p/flipside/cardservice/api/transfer"
	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/service/training"
)

var (
	ErrUnresolvedTrainingType = errors.New("Unresolver training type")
)

type TrainingController struct {
	service *training.TrainingService
}

func NewTrainingController(service *training.TrainingService) *TrainingController {
	return &TrainingController{
		service: service,
	}
}

func (c *TrainingController) GetCardSetTrainings(userId int, slug string) ([]transfer.TrainingSummaryResponse, error) {
	trainings, err := c.service.GetCardSetTrainings(userId, slug)
	if err != nil {
		return nil, err
	}
	response := make([]transfer.TrainingSummaryResponse, len(trainings))
	for i := range trainings {
		response[i] = transfer.TrainingSummaryResponse{
			Id:           trainings[i].Id,
			CardSetId:    trainings[i].CardSetId,
			TrainingType: string(trainings[i].TrainingType),
			Status:       trainings[i].Status,
			CountRight:   trainings[i].CountRight,
			CountWrong:   trainings[i].CountWrong,
		}
	}
	return response, nil
}

func (c *TrainingController) GetTrainingSummary(userId int, trainingId string) (*transfer.TrainingSummaryResponse, error) {
	trainingIdInt, err := strconv.Atoi(trainingId)
	if err != nil {
		return nil, err
	}
	res, err := c.service.GetTrainingSummary(userId, trainingIdInt)
	if err != nil {
		return nil, err
	}
	return &transfer.TrainingSummaryResponse{
		Id:           res.Id,
		TrainingType: string(res.TrainingType),
		Status:       res.Status,
		CountRight:   res.CountRight,
		CountWrong:   res.CountWrong,
	}, nil
}

func (c *TrainingController) CreateTraining(userId int, slug string, trainingType string) (*transfer.TrainingSummaryResponse, error) {
	trainingTypeModel, err := resolveTrainingType(trainingType)
	if err != nil {
		return nil, err
	}
	res, err := c.service.CreateTraining(userId, slug, *trainingTypeModel)
	if err != nil {
		return nil, err
	}
	return &transfer.TrainingSummaryResponse{
		Id:           res.Id,
		CardSetId:    res.CardSetId,
		TrainingType: string(res.TrainingType),
		Status:       res.Status,
		CountRight:   res.CountRight,
		CountWrong:   res.CountWrong,
	}, nil
}

func (c *TrainingController) GetNextTask(userId int, trainingId string) (*transfer.TaskResponse, error) {
	trainingIdInt, err := strconv.Atoi(trainingId)
	if err != nil {
		return nil, err
	}
	res, err := c.service.GetNextTask(userId, trainingIdInt)
	if err != nil {
		return nil, err
	}
	return &transfer.TaskResponse{
		Question:     res.Question,
		QuestionType: string(res.QuestionType),
		Answers:      res.Answers,
	}, nil
}

func (c *TrainingController) SubmitTask(userId int, trainingId string, answer string) (*transfer.TaskResultResponse, error) {
	trainingIdInt, err := strconv.Atoi(trainingId)
	if err != nil {
		return nil, err
	}
	res, err := c.service.SubmitAnswer(userId, trainingIdInt, answer)
	if err != nil {
		return nil, err
	}
	return &transfer.TaskResultResponse{
		Answer:    *res.Answer,
		IsCorrect: res.IsCorrect,
	}, nil
}

func resolveTrainingType(trainingTypeStr string) (*model.TrainingType, error) {
	var res model.TrainingType
	switch model.TrainingType(trainingTypeStr) {
	case model.TrainingTypeBasic:
		res = model.TrainingTypeBasic
	case model.TrainingTypeQuiz:
		res = model.TrainingTypeQuiz
	default:
		return nil, ErrUnresolvedTrainingType
	}
	return &res, nil
}
