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
	Id           int          `db:"id"`
	UserId       int          `db:"user_id"`
	CardSetId    int          `db:"card_set_id"`
	TrainingType TrainingType `db:"training_type"`
	Status       string       `db:"status"`
	CreatedAt    time.Time    `db:"created_at"`
	FinishedAt   *time.Time   `db:"finished_at"`
}

type TrainingTaskResult struct {
	Id            int       `db:"id"`
	TrainingId    int       `db:"training_id"`
	CardId        int       `db:"card_id"`
	Answer        *string   `db:"answer"`
	CorrectAnswer *string   `db:"correct_answer"`
	IsCorrect     bool      `db:"is_correct"`
	CreatedAt     time.Time `db:"created_at"`
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
