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
	// GetNextTask returns next task and сan save it to trainingRepository
	GetNextTask(training *model.Training) (*CheckerTask, error)

	// Validate validates answer
	Validate(training *model.Training, answer string) (bool, error)

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
	lastTask, err := c.trainingRepository.GetLastTaskResult(training.Id)
	if err != nil && !errors.Is(err, repository.ErrTrainingTaskResultNotFound) {
		return nil, err
	}
	if lastTask != nil && lastTask.Answer == nil {
		return &CheckerTask{
			TrainingId:   training.Id,
			CardId:       lastTask.CardId,
			QuestionType: model.QuestionTypeBinary,
			Answers:      []string{knowAnswer, dontKnowAnswer},
		}, nil
	}
	doneCards, err := c.trainingRepository.GetTaskResults(training.Id)
	if err != nil {
		return nil, err
	}
	idsSet := make(map[int]struct{}, len(doneCards))
	for i := range doneCards {
		idsSet[doneCards[i].CardId] = struct{}{}
	}
	cards, err := c.cardRepository.GetCardsByCardSet(training.CardSetId)
	var task *CheckerTask
	for _, i := range rand.Perm(len(cards)) {
		if _, found := idsSet[cards[i].Id]; !found {
			task = &CheckerTask{
				TrainingId:   training.Id,
				CardId:       cards[i].Id,
				QuestionType: model.QuestionTypeBinary,
				Answers:      []string{cards[i].Answer},
			}
			break
		}
	}
	if task == nil {
		return nil, ErrAllTasksAreCompleted
	}
	taskResult := model.TrainingTaskResult{
		TrainingId: training.Id,
		CardId:     task.CardId,
	}
	if _, err := c.trainingRepository.CreateTaskResult(&taskResult); err != nil {
		return nil, err
	}
	return task, nil
}

func (c *BasicTaskChecker) Validate(training *model.Training, answer string) (bool, error) {
	lastTask, err := c.trainingRepository.GetLastTaskResult(training.Id)
	if err != nil {
		return false, err
	}
	if lastTask == nil || lastTask.Answer != nil {
		return false, ErrTaskNotFound
	}
	if answer == knowAnswer {
		return true, nil
	} else if answer == dontKnowAnswer {
		return false, nil
	} else {
		return false, ErrInvalidAnswer
	}
}

func (c *BasicTaskChecker) IsSupporting(trainingType model.TrainingType) bool {
	return trainingType == model.TrainingTypeBasic
}
