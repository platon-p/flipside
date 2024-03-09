package transfer

type TrainingSummaryResponse struct {
	Id           int    `json:"id"`
	TrainingType string `json:"training_type"`
	Status       string `json:"status"`
	CountRight   int    `json:"count_right"`
	CountWrong   int    `json:"count_wrong"`
}

type TaskResponse struct {
	Question     string   `json:"question"`
	QuestionType string   `json:"question_type"`
	Answers      []string `json:"answers"`
}

type TaskResultResponse struct {
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	IsCorrect bool   `json:"is_correct"`
}
