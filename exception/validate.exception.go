package exception

type ValidateException struct {
	Message string
}

func NewValidateException(message string) ValidateException {
	return ValidateException{Message: message}
}

func (exception ValidateException) Error() string {
	return exception.Message
}
