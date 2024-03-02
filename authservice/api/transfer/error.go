package transfer

type ErrorResponse struct {
    StatusCode int `json:"status_code"`
    Message string `json:"message"`
}
