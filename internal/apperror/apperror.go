package apperror

type AppError struct {
	Err     error  `json:"-"`
	Type    string `json:"type,omitempty"`
	Message string `json:"error,omitempty"`
}

func New(e error, t string, msg string) *AppError {
	return &AppError{Err: e, Message: msg, Type: t}
}

func NewEntityError(e error, msg string) *AppError {
	return &AppError{Type: "entity", Message: msg}
}

func NewRepoError(e error, msg string) *AppError {
	return &AppError{Type: "repository", Message: msg}
}

func NewHandlerError(e error, msg string) *AppError {
	return &AppError{Type: "handler", Message: msg}
}

func NewServiceError(e error, msg string) *AppError {
	return &AppError{Type: "service", Message: msg}
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}
