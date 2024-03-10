package training

import (
	"errors"
	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/repository"
	"math/rand"
)

var (
	ErrAllTasksAreCompleted = errors.New("All tasks are completed")
	ErrInvalidAnswer        = errors.New("Invalid answer")
	ErrTaskNotFound         = errors.New("Task not found")

	knowAnswer     = "Know"
	dontKnowAnswer = "Don't know"
)

type CheckerTask struct {
	TrainingId   int
	CardId       int
	QuestionType model.QuestionType
	Answers      []string
}

type TaskChecker interface {
	GetNextTask(training *model.Training) (*CheckerTask, error)

	// Does not save result to trainingRepository
	Validate(card *model.Card, answer string) (bool, error)

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

func (c *BasicTaskChecker) GetNextTask(training *model.Training) (*CheckerTask, error) {
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
			return &CheckerTask{
				TrainingId:   training.Id,
				CardId:       cards[i].Id,
				QuestionType: model.QuestionTypeBinary,
				Answers:      answer,
			}, nil
		}
	}
	return nil, ErrAllTasksAreCompleted
}

func (c *BasicTaskChecker) Validate(card *model.Card, answer string) (bool, error) {
	if answer == knowAnswer {
		return true, nil
	} else if answer == dontKnowAnswer {
		return true, nil
	} else {
		return false, ErrInvalidAnswer
	}
}

func (c *BasicTaskChecker) IsSupporting(trainingType model.TrainingType) bool {
	return trainingType == model.TrainingTypeBasic
}
