package deftype

import (
	"net/http"

	"go.uber.org/zap"
)

type Error interface {
	Number() int64
	Status() int
	Error() string
}

type CustomError struct {
	number     int64
	statusCode int
	message    string
}

func (r CustomError) Number() int64 {
	return r.number
}

func (r CustomError) Status() int {
	return r.statusCode
}

func (r CustomError) Error() string {
	return r.message
}

var errs = map[int64]Error{}

func setErrs(num int64, code int, msg string) {
	if _, duplicated := errs[num]; duplicated {
		zap.S().Panic("duplicated error!!")
		return
	}

	errs[num] = New(num, code, msg)
}

func New(num int64, code int, msg string) Error {
	setErrs(num, code, msg)

	return CustomError{
		number:     num,
		message:    msg,
		statusCode: code,
	}
}

var (
	ErrInvalidRequestData = New(1, http.StatusBadRequest, "invalid request data")

	ErrInternalServerError = New(5000, http.StatusInternalServerError, "internal server error")
)
