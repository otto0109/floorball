package api

type ResponseError struct {
	Code  int
	Error error
}
