package model

import "time"

const (
	TrainingTypeBasic TrainingType = "basic"
	TrainingTypeQuiz  TrainingType = "quiz"
)

type TrainingType string

type Training struct {
	Id           int
	UserId       int
	CardSetId    int
	TrainingType TrainingType
	Status       string
	CreatedAt    time.Time
	FinishedAt   time.Time
}

type TrainingTaskResult struct {
	Id         int
	TrainingId int
	CardId     int
	Answer     string
	IsCorrect  bool
	CreatedAt  time.Time
}

type TrainingSummary struct {
	Id           int
	Status       string
	TrainingType TrainingType
	CountRight   int
	CountWrong   int
}

type Task struct {
	Question string
	TaskType string
	Answers  []string
}
