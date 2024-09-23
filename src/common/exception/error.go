package exception

type NotFoundError struct {
	Error string
}

type GeneralError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func NewGeneralError(error string) GeneralError {
	return GeneralError{Error: error}
}
