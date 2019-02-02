package users

const (
	successStatus = "success"
	failureStatus = "failed"
)

type response struct {
	Status    string `json:"status"`
	Error     string `json:"error,omitempty"`
	RequestID string `json:"request_id"`
	responseData
}

type responseData struct {
	User *User `json:"user,omitempty"`
}

