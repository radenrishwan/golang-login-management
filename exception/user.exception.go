package exception

type UserException struct {
	Message string
}

func NewUserException(message string) ValidateException {
	return ValidateException{Message: message}
}

func (exception UserException) Error() string {
	return exception.Message
}
