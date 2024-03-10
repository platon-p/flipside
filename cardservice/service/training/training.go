package training

import (
	"errors"

	"fmt"

	"github.com/platon-p/flipside/cardservice/model"
	"github.com/platon-p/flipside/cardservice/repository"
	"github.com/platon-p/flipside/cardservice/service"
)

var (
	ErrNotATrainingOwner   = errors.New("Not a training owner")
	ErrTrainingIsCompleted = errors.New("Training is completed")
)

type TrainingService struct {
	trainingRepository repository.TrainingRepository
	cardSetRepository  repository.CardSetRepository

	checkers []TaskChecker
}

func NewTrainingService(trainingRepository repository.TrainingRepository, cardSetRepository repository.CardSetRepository, checkers []TaskChecker) *TrainingService {
	return &TrainingService{
		trainingRepository: trainingRepository,
		cardSetRepository:  cardSetRepository,
		checkers:           checkers,
	}
}

func (s *TrainingService) GetCardSetTrainings(userId int, slug string) ([]model.TrainingSummary, error) {
	cardSet, err := s.cardSetRepository.GetCardSet(slug)
	if errors.Is(err, repository.ErrCardSetNotFound) {
		return nil, service.ErrCardSetNotFound
	}
	trainings, err := s.trainingRepository.GetCardSetTrainings(userId, cardSet.Id)
	if err != nil {
		return nil, err
	}
	summaries := make([]model.TrainingSummary, len(trainings))
	for i := range trainings {
		summary, err := s.MakeTrainingSummary(&trainings[i])
		if err != nil {
			return nil, err
		}
		summaries[i] = *summary
	}
	return summaries, nil
}

func (s *TrainingService) CreateTraining(userId int, slug string, trainingType model.TrainingType) (*model.TrainingSummary, error) {
	cardSet, err := s.cardSetRepository.GetCardSet(slug)
	if errors.Is(err, repository.ErrCardSetNotFound) {
		return nil, service.ErrCardSetNotFound
	}
	training := model.Training{
		UserId:       userId,
		CardSetId:    cardSet.Id,
		TrainingType: trainingType,
		Status:       model.TrainingStatusCreated,
	}
	newEntity, err := s.trainingRepository.CreateTraining(&training)
	if err != nil {
		return nil, err
	}
	return s.MakeTrainingSummary(newEntity)
}

func (s *TrainingService) GetTrainingSummary(userId int, trainingId int) (*model.TrainingSummary, error) {
	training, err := s.GetTraining(userId, trainingId)
	if err != nil {
		return nil, err
	}
	return s.MakeTrainingSummary(training)
}

func (s *TrainingService) MakeTrainingSummary(training *model.Training) (*model.TrainingSummary, error) {
	results, err := s.trainingRepository.GetTaskResults(training.Id)
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
		CardSetId:    training.CardSetId,
		Status:       training.Status,
		TrainingType: training.TrainingType,
		CreatedAt:    training.CreatedAt,
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
	if training.Status == model.TrainingStatusCreated {
		if _, err := s.trainingRepository.SetTrainingStatus(trainingId, model.TrainingStatusStarted); err != nil {
			return nil, err
		}
	}
	checker := s.resolveChecker(training.TrainingType)
	task, err := checker.GetNextTask(training)
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

func (s *TrainingService) SubmitAnswer(userId int, trainingId int, answer string) (*model.TrainingTaskResult, error) {
	training, err := s.GetTraining(userId, trainingId)
	if err != nil {
		return nil, err
	}
	if training.Status == model.TrainingStatusCompleted {
		return nil, ErrTrainingIsCompleted
	}
	result, err := s.resolveChecker(training.TrainingType).Submit(training, answer)
	if err != nil {
		return nil, err
	}
	res, err := s.trainingRepository.CreateTaskResult(trainingId, result)
	if err != nil {
		return nil, err
	}
	return res, nil
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
