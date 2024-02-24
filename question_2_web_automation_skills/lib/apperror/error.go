package apperror

const (
	PageInvalid = iota + 1
	ElementInvalid
	ResultInvalid
)

type CommonError struct {
	ParentError error  `json:"-"`
	Code        int    `json:"code"`
	Message     string `json:"message"`
}

func (e CommonError) Error() string {
	return e.Message
}

func NewError(err error, message string, code int) error {
	return CommonError{ParentError: err, Message: message, Code: code}
}

func NewErrorBase(message string, code int) error {
	return CommonError{Message: message, Code: code}
}
