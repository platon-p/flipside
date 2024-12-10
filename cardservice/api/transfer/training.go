package transfer

import (
	"errors"

	"github.com/platon-p/flipside/cardservice/model"
)

var (
	ErrUnresolvedTrainingType = errors.New("Unresolver training type")
)

type TrainingSummaryResponse struct {
	Id           int    `json:"id"`
	CardSetId    int    `json:"card_set_id"`
	TrainingType string `json:"training_type"`
	Status       string `json:"status"`
	CountRight   int    `json:"count_right"`
	CountWrong   int    `json:"count_wrong"`
}

type TaskResponse struct {
	Question     string   `json:"question"`
	QuestionType string   `json:"question_type"`
	Answers      []string `json:"answers"`
}

type TaskResultResponse struct {
	Answer    string `json:"answer"`
	IsCorrect bool   `json:"is_correct"`
}

func TrainingSummaryToResponse(mdl model.TrainingSummary) TrainingSummaryResponse {
	return TrainingSummaryResponse{
		Id:           mdl.Id,
		CardSetId:    mdl.CardSetId,
		TrainingType: string(mdl.TrainingType),
		Status:       mdl.Status,
		CountRight:   mdl.CountRight,
		CountWrong:   mdl.CountWrong,
	}
}

func TrainingSummariesToResponse(trainings []model.TrainingSummary) []TrainingSummaryResponse {
	response := make([]TrainingSummaryResponse, len(trainings))
	for i := range trainings {
		response[i] = TrainingSummaryResponse{
			Id:           trainings[i].Id,
			CardSetId:    trainings[i].CardSetId,
			TrainingType: string(trainings[i].TrainingType),
			Status:       trainings[i].Status,
			CountRight:   trainings[i].CountRight,
			CountWrong:   trainings[i].CountWrong,
		}
	}
	return response
}

func ResolveTrainingType(trainingTypeStr string) (model.TrainingType, error) {
	switch model.TrainingType(trainingTypeStr) {
	case model.TrainingTypeBasic:
		return model.TrainingTypeBasic, nil
	case model.TrainingTypeQuiz:
		return model.TrainingTypeQuiz, nil
	default:
		return "", ErrUnresolvedTrainingType
	}
}

func TaskToResponse(task *model.Task) TaskResponse{
    return TaskResponse{
		Question:     task.Question,
		QuestionType: string(task.QuestionType),
		Answers:      task.Answers,
	}
}
