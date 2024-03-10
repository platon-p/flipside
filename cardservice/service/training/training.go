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
	ErrTrainingNotFound    = errors.New("Training not found")
)

type TrainingService struct {
	trainingRepository repository.TrainingRepository
	cardSetRepository  repository.CardSetRepository
	cardRepository     repository.CardRepository

	checkers []TaskChecker
}

func NewTrainingService(trainingRepository repository.TrainingRepository, cardSetRepository repository.CardSetRepository, cardRepository repository.CardRepository, checkers []TaskChecker) *TrainingService {
	return &TrainingService{
		trainingRepository: trainingRepository,
		cardSetRepository:  cardSetRepository,
		cardRepository:     cardRepository,
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
		summary, err := s.makeTrainingSummary(&trainings[i])
		if err != nil {
			return nil, err
		}
		summaries[i] = *summary
	}
	return summaries, nil
}

func (s *TrainingService) GetTrainingSummary(userId int, trainingId int) (*model.TrainingSummary, error) {
	training, err := s.getTraining(userId, trainingId)
	if err != nil {
		return nil, err
	}
	return s.makeTrainingSummary(training)
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
	return s.makeTrainingSummary(newEntity)
}

func (s *TrainingService) GetNextTask(userId int, trainingId int) (*model.Task, error) {
	training, err := s.getTraining(userId, trainingId)
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
	card, err := s.cardRepository.GetCard(task.CardId)
	if err != nil {
		return nil, err
	}
	taskResponse := model.Task{
		Question:     card.Question,
		QuestionType: task.QuestionType,
		Answers:      task.Answers,
	}
	return &taskResponse, nil
}

func (s *TrainingService) SubmitAnswer(userId int, trainingId int, answer string) (*model.TrainingTaskResult, error) {
	training, err := s.getTraining(userId, trainingId)
	if err != nil {
		return nil, err
	}
	if training.Status == model.TrainingStatusCompleted {
		return nil, ErrTrainingIsCompleted
	}
	lastTask, err := s.trainingRepository.GetLastTaskResult(trainingId)
	if errors.Is(err, repository.ErrTrainingTaskResultNotFound) {
		return nil, ErrTaskNotFound
	}
	isCorrect, err := s.resolveChecker(training.TrainingType).Validate(training, answer)
	if err != nil {
		return nil, err
	}
	lastTask.Answer = &answer
	lastTask.IsCorrect = &isCorrect
	res, err := s.trainingRepository.SaveTaskResult(lastTask)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TrainingService) getTraining(userId int, trainingId int) (*model.Training, error) {
	training, err := s.trainingRepository.GetTraining(trainingId)
	if errors.Is(err, repository.ErrTrainingNotFound) {
		return nil, ErrTrainingNotFound
	}
	if err != nil {
		return nil, err
	}
	if training.UserId != userId {
		return nil, ErrNotATrainingOwner
	}
	return training, nil
}

func (s *TrainingService) makeTrainingSummary(training *model.Training) (*model.TrainingSummary, error) {
	results, err := s.trainingRepository.GetTaskResults(training.Id)
	if err != nil {
		return nil, err
	}
	correct, wrong := 0, 0
	for i := range results {
		if results[i].IsCorrect == nil {
			continue
		}
		if *results[i].IsCorrect {
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

func (s *TrainingService) resolveChecker(trainingType model.TrainingType) TaskChecker {
	for i := range s.checkers {
		if s.checkers[i].IsSupporting(trainingType) {
			return s.checkers[i]
		}
	}
	fmt.Println("Checker not found")
	return nil
}
