package training

import (
	"errors"
	"math/rand"

	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/repository"
)

var (
	ErrAllTasksAreCompleted = errors.New("All tasks are completed")
	ErrInvalidAnswer        = errors.New("Invalid answer")

	knowAnswer     = "Know"
	dontKnowAnswer = "Don't know"
)

type TaskChecker interface {
	GetNextTask(trainingId int) (*model.Task, error)

	// Does not save result to repository
	Submit(training *model.Training, answer string) (*model.TrainingTaskResult, error)

	IsSupporting(trainingType model.TrainingType) bool
}

type BasicTaskChecker struct {
	repository     repository.TrainingRepository
	cardRepository repository.CardRepository
}

func (c *BasicTaskChecker) IsSupporting(trainingType model.TrainingType) bool {
	return trainingType == model.TrainingTypeBasic
}

func (c *BasicTaskChecker) GetNextTask(trainingId int, cardSetId int) (*model.Task, error) {
	doneCards, err := c.repository.GetTaskResults(trainingId)
	if err != nil {
		return nil, err
	}
	idsSet := make(map[int]struct{}, len(doneCards))
	for i := range doneCards {
		idsSet[doneCards[i].CardId] = struct{}{}
	}
	cards, err := c.cardRepository.GetCardsByCardSet(cardSetId)
	answer := []string{knowAnswer, dontKnowAnswer}
	for _, v := range rand.Perm(len(cards)) {
		if _, found := idsSet[cards[v].Id]; !found {
			return &model.Task{
				Question:     cards[v].Question,
				TrainingType: model.TrainingTypeBasic,
				Answers:      answer,
			}, nil
		}
	}
	return nil, ErrAllTasksAreCompleted
}

func (c *BasicTaskChecker) Submit(training *model.Training, answer string) (*model.TrainingTaskResult, error) {
	results, err := c.repository.GetTaskResults(training.Id)
	if err != nil {
		return nil, err
	}
	task := results[len(results)-1]
	isCorrect := false
	if answer == knowAnswer {
		isCorrect = true
	} else if answer != dontKnowAnswer {
		return nil, ErrInvalidAnswer
	}
	result := model.TrainingTaskResult{
		TrainingId: training.Id,
		CardId:     task.CardId,
		Answer:     &answer,
		IsCorrect:  isCorrect,
	}
    return &result, nil
}
