package v1

// Форма возврата
type UploadResponse struct {
	Msg string `json:"message"`
}

// Response при order create
type IdResponse struct {
	Id string `json:"id"`
}
