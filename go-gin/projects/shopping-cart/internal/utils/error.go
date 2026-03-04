package utils

type ErrorCode string

const (
	ErrCodeBadRequest          ErrorCode = "BAD_REQUEST"
	ErrCodeNotFound            ErrorCode = "NOT_FOUND"
	ErrCodeInternalServerError ErrorCode = "INTERNAL_SERVER_ERROR"
	ErrCodeConflict            ErrorCode = "CONFLICT"
)

type AppError struct {
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`
	Err     error     `json:"error,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewError(message string, code ErrorCode) error {
	return &AppError{
		Message: message,
		Code:    code,
	}
}

func WrapError(message string, code ErrorCode, err error) error {
	return &AppError{
		Message: message,
		Code:    code,
		Err:     err,
	}
}
