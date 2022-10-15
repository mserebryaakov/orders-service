package apperror

var (
	ErrNotFound = NewAppError(nil, "Not found")
)

type AppError struct {
	Err     error  `json:"-"`
	Message string `json:"message,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewAppError(err error, message string) *AppError {
	return &AppError{
		Err:     err,
		Message: message,
	}
}
