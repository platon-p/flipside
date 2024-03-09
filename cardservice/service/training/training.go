package training

import (
	"errors"

	"fmt"

	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/repository"
	"github.com/platon-p/flipside/cardservice/service"
)

var (
	ErrNotATrainingOwner = errors.New("Not a training owner")
    ErrTrainingIsCompleted = errors.New("Training is completed")
)

type TrainingService struct {
	trainingRepository repository.TrainingRepository
	cardSetRepository  repository.CardSetRepository

	checkers []TaskChecker
}

func (s *TrainingService) CreateTraining(userId int, slug string, trainingType model.TrainingType) (*model.Training, error) {
	cardSet, err := s.cardSetRepository.GetCardSet(slug)
	if errors.Is(err, repository.ErrCardSetNotFound) {
		return nil, service.ErrCardSetNotFound
	}
	training := model.Training{
		UserId:       userId,
		CardSetId:    cardSet.Id,
		TrainingType: trainingType,
		Status:       "Created",
	}
	return s.trainingRepository.CreateTraining(&training)
}

func (s *TrainingService) GetTrainingSummary(userId int, trainingId int) (*model.TrainingSummary, error) {
	training, err := s.GetTraining(userId, trainingId)
	if err != nil {
		return nil, err
	}
    if training.Status == model.TrainingStatusCompleted {
        return nil, ErrTrainingIsCompleted
    }
	results, err := s.trainingRepository.GetTaskResults(trainingId)
	if err != nil {
		return nil, err
	}
	correct, wrong := 0, 0
	for i := range results {
		if results[i].IsCorrect {
			correct++
		} else {
			wrong++
		}
	}
	summary := model.TrainingSummary{
		Id:           training.Id,
		Status:       training.Status,
		TrainingType: training.TrainingType,
		CountRight:   correct,
		CountWrong:   wrong,
	}
	return &summary, nil
}

func (s *TrainingService) GetNextTask(userId int, trainingId int) (*model.Task, error) {
	training, err := s.GetTraining(userId, trainingId)
	if err != nil {
		return nil, err
	}
    if training.Status == model.TrainingStatusCompleted {
        return nil, ErrTrainingIsCompleted
    }
	checker := s.resolveChecker(training.TrainingType)
    task, err := checker.GetNextTask(trainingId)
    if errors.Is(err, ErrAllTasksAreCompleted) {
        if _, err := s.trainingRepository.SetTrainingStatus(training.Id, model.TrainingStatusCompleted); err != nil {
            return nil, err
        }
        return nil, ErrTrainingIsCompleted
    }
    if err != nil {
        return nil, err
    }
    return task, nil
    
}

func (s *TrainingService) SubmitAnswer(userId int, trainingId int, answer string) error {
	training, err := s.GetTraining(userId, trainingId)
	if err != nil {
		return err
	}
    if training.Status == model.TrainingStatusCompleted {
        return ErrTrainingIsCompleted
    }
	result, err := s.resolveChecker(training.TrainingType).Submit(training, answer)
	if err != nil {
		return err
	}
	_, err = s.trainingRepository.CreateTaskResult(trainingId, result)
	return err
}

func (s *TrainingService) GetTraining(userId int, trainingId int) (*model.Training, error) {
	training, err := s.trainingRepository.GetTraining(trainingId)
	if err != nil {
		return nil, err
	}
	if training.UserId != userId {
		return nil, ErrNotATrainingOwner
	}
	return training, nil
}

func (s *TrainingService) resolveChecker(trainingType model.TrainingType) TaskChecker {
	for i := range s.checkers {
		if s.checkers[i].IsSupporting(trainingType) {
			return s.checkers[i]
		}
	}
	fmt.Println("Checker not found")
	return nil
}
