package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/platon-p/flipside/cardservice/model"
)

var (
	trainingsTable            = "trainings"
	trainingTasksResultsTable = "training_tasks_results"

	ErrTrainingNotFound           = errors.New("Training not found")
	ErrTrainingTaskResultNotFound = errors.New("Training task result not found")
)

type TrainingRepository interface {
	GetTraining(trainingId int) (*model.Training, error)
	CreateTraining(training *model.Training) (*model.Training, error)
	GetCardSetTrainings(userId int, cardSetId int) ([]model.Training, error)
	SetTrainingStatus(trainingId int, status string) (*model.Training, error)

	GetTaskResults(trainingId int) ([]model.TrainingTaskResult, error)
	GetLastTaskResult(trainingId int) (*model.TrainingTaskResult, error)
	CreateTaskResult(trainingId int, taskResult *model.TrainingTaskResult) (*model.TrainingTaskResult, error)
}

type TrainingRepositoryImpl struct {
	db *sqlx.DB
}

func NewTrainingRepositoryImpl(db *sqlx.DB) *TrainingRepositoryImpl {
	return &TrainingRepositoryImpl{db: db}
}

func (r *TrainingRepositoryImpl) GetTraining(trainingId int) (*model.Training, error) {
	query := fmt.Sprintf(`SELECT * FROM %v WHERE id = $1`, trainingsTable)
	var found model.Training
	err := r.db.QueryRowx(query, trainingId).StructScan(&found)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrTrainingNotFound
	}
	if err != nil {
		return nil, err
	}
	return &found, nil
}

func (r *TrainingRepositoryImpl) CreateTraining(training *model.Training) (*model.Training, error) {
	query := fmt.Sprintf(
		`INSERT INTO %v(user_id, card_set_id, training_type, status)
        VALUES ($1, $2, $3, $4)
        RETURNING *`,
		trainingsTable,
	)
	var newEntity model.Training
	err := r.db.QueryRowx(query, training.UserId, training.CardSetId, training.TrainingType, training.Status).
		StructScan(&newEntity)
	if err != nil {
		return nil, err
	}
	return &newEntity, nil
}

func (r *TrainingRepositoryImpl) GetCardSetTrainings(userId int, cardSetId int) ([]model.Training, error) {
	query := fmt.Sprintf(`SELECT * FROM %v WHERE user_id = $1 AND card_set_id = $2`, trainingsTable)
	rows, err := r.db.Queryx(query, userId, cardSetId)
	if err != nil {
		return nil, err
	}
	trainings := make([]model.Training, 0)
	for rows.Next() {
		var row model.Training
		if err := rows.StructScan(&row); err != nil {
			return nil, err
		}
		trainings = append(trainings, row)
	}
	return trainings, nil
}

func (r *TrainingRepositoryImpl) SetTrainingStatus(trainingId int, status string) (*model.Training, error) {
	query := fmt.Sprintf(`UPDATE %v SET status = $1 WHERE id = $2`, trainingsTable)
	var updated model.Training
	err := r.db.QueryRowx(query, status, trainingId).StructScan(&updated)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *TrainingRepositoryImpl) GetTaskResults(trainingId int) ([]model.TrainingTaskResult, error) {
	query := fmt.Sprintf(`SELECT * FROM %v WHERE training_id = $1`, trainingTasksResultsTable)
	rows, err := r.db.Queryx(query, trainingId)
	if err != nil {
		return nil, err
	}
	results := make([]model.TrainingTaskResult, 0)
	for rows.Next() {
		var row model.TrainingTaskResult
		if err := rows.StructScan(&row); err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil
}

func (r *TrainingRepositoryImpl) GetLastTaskResult(trainingId int) (*model.TrainingTaskResult, error) {
	query := fmt.Sprintf(
		`SELECT * FROM %v WHERE training_id = $1
                 ORDER BY created_at DESC LIMIT 1`,
		trainingTasksResultsTable,
	)
	var found model.TrainingTaskResult
	err := r.db.QueryRowx(query, trainingId).StructScan(&found)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrTrainingTaskResultNotFound
	}
	if err != nil {
		return nil, err
	}
	return &found, nil
}

func (r *TrainingRepositoryImpl) CreateTaskResult(trainingId int, taskResult *model.TrainingTaskResult) (*model.TrainingTaskResult, error) {
	query := fmt.Sprintf(
		`INSERT INTO %v(training_id, card_id, answer, correct_answer, is_correct) 
        VALUES ($1, $2, $3, $4, $5)
        RETURNING *;`,
		trainingTasksResultsTable,
	)
	var newEntity model.TrainingTaskResult
	err := r.db.
		QueryRowx(query, trainingId, taskResult.CardId, taskResult.Answer, taskResult.CorrectAnswer, taskResult.IsCorrect).
		StructScan(&newEntity)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("asdasdasd", err)
	}
	if err != nil {
		return nil, err
	}
	return &newEntity, nil
}
