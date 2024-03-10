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
	ErrTaskNotFound         = errors.New("Task not found")

	knowAnswer     = "Know"
	dontKnowAnswer = "Don't know"
)

type TaskChecker interface {
	GetNextTask(training *model.Training) (*model.Task, error)

	// Does not save result to trainingRepository
	Submit(training *model.Training, answer string) (*model.TrainingTaskResult, error)

	IsSupporting(trainingType model.TrainingType) bool
}

type BasicTaskChecker struct {
	trainingRepository repository.TrainingRepository
	cardRepository     repository.CardRepository
}

func NewBasicTaskChecker(repository repository.TrainingRepository, cardRepository repository.CardRepository) *BasicTaskChecker {
	return &BasicTaskChecker{
		trainingRepository: repository,
		cardRepository:     cardRepository,
	}
}

func (c *BasicTaskChecker) GetNextTask(training *model.Training) (*model.Task, error) {
	doneCards, err := c.trainingRepository.GetTaskResults(training.Id)
	if err != nil {
		return nil, err
	}
	idsSet := make(map[int]struct{}, len(doneCards))
	for i := range doneCards {
		idsSet[doneCards[i].CardId] = struct{}{}
	}
	cards, err := c.cardRepository.GetCardsByCardSet(training.CardSetId)
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
	lastQuestion, err := c.trainingRepository.GetLastTaskResult(training.Id)
	if errors.Is(err, repository.ErrTrainingTaskResultNotFound) {
		return nil, ErrTaskNotFound
	}
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

func (c *BasicTaskChecker) IsSupporting(trainingType model.TrainingType) bool {
	return trainingType == model.TrainingTypeBasic
}
