package customError

func (err *BadRequest) Error() string {
	return err.ErrorText
}

type BadRequest struct {
	ErrorText string
}
