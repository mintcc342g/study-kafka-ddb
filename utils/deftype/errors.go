package deftype

import (
	"net/http"
	"sync"

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
var mu = sync.RWMutex{}

func checkDup(num int64, code int, msg string) {
	if _, duplicated := errs[num]; duplicated {
		zap.S().Panic("duplicated error!!")
		return
	}
}

func New(num int64, code int, msg string) Error {
	checkDup(num, code, msg)

	err := CustomError{
		number:     num,
		message:    msg,
		statusCode: code,
	}

	mu.Lock()
	errs[num] = err
	mu.Unlock()

	return err
}

var (
	ErrInvalidRequestData = New(1, http.StatusBadRequest, "invalid request data")
	ErrUnauthorized       = New(2, http.StatusUnauthorized, "the user has no permission")

	ErrInternalServerError = New(5000, http.StatusInternalServerError, "internal server error")
)
