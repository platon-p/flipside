package training

import (
	"errors"
	"math/rand"
	"time"

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
	for _, i := range rand.Perm(len(cards)) {
		if _, found := idsSet[cards[i].Id]; !found {
			return &model.Task{
				Question:     cards[i].Question,
				QuestionType: model.QuestionTypeBinary,
				Answers:      answer,
			}, nil
		}
	}
	return nil, ErrAllTasksAreCompleted
}

func (c *BasicTaskChecker) Submit(training *model.Training, answer string) (*model.TrainingTaskResult, error) {
	lastQuestion, err := c.repository.GetLastTaskResult(training.Id)
	if err != nil {
		return nil, err
	}
	card, err := c.cardRepository.GetCard(lastQuestion.CardId)
	isCorrect := false
	if answer == knowAnswer {
		isCorrect = true
	} else if answer != dontKnowAnswer {
		return nil, ErrInvalidAnswer
	}
	result := model.TrainingTaskResult{
		TrainingId:    lastQuestion.TrainingId,
		CardId:        lastQuestion.CardId,
		Answer:        &answer,
		CorrectAnswer: &card.Answer,
		IsCorrect:     isCorrect,
		CreatedAt:     time.Now(),
	}
	return &result, nil
}
