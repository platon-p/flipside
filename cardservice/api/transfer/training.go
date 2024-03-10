package transfer

type TrainingSummaryResponse struct {
	Id           int    `json:"id"`
	CardSetId    int    `json:"card_set_id"`
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
	Answer    string `json:"answer"`
	IsCorrect bool   `json:"is_correct"`
}
