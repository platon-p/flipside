package model

import "time"

const (
	TrainingTypeBasic TrainingType = "basic"
	TrainingTypeQuiz  TrainingType = "quiz"

	TrainingStatusCompleted = "COMPLETED"
	TrainingStatusCreated   = "CREATED"
	TrainingStatusStarted   = "STARTED"

	QuestionTypeBinary QuestionType = "Binary"
	QuestionTypeChoice QuestionType = "Choice"
	QuestionTypeInput  QuestionType = "Input"
)

type QuestionType string

type TrainingType string

type Training struct {
	Id           int
	UserId       int
	CardSetId    int
	TrainingType TrainingType
	Status       string
	CreatedAt    time.Time
	FinishedAt   *time.Time
}

type TrainingTaskResult struct {
	Id            int
	TrainingId    int
	CardId        int
	Answer        *string
	CorrectAnswer *string
	IsCorrect     bool
	CreatedAt     time.Time
}

type TrainingSummary struct {
	Id           int
	CardSetId    int
	Status       string
	TrainingType TrainingType
	CreatedAt    time.Time

	CountRight int
	CountWrong int
}

type Task struct {
	Question     string
	QuestionType QuestionType
	Answers      []string
}
