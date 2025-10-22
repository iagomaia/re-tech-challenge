package utils

import (
	"fmt"
	"net/http"
	"reflect"
)

type CustomError struct {
	Status        int
	Message       string
	OriginalError error
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%v: %v", e.Message, e.OriginalError)
}

func GetCustomError(err error) CustomError {
	cErr := CustomError{
		Status:        http.StatusInternalServerError,
		Message:       "Unexpected error happened",
		OriginalError: err,
	}

	if reflect.TypeOf(err).Name() == reflect.TypeOf(cErr).Name() {
		metaValue := reflect.ValueOf(err)
		cErr.Status = int(metaValue.FieldByName("Status").Int())
		cErr.Message = metaValue.FieldByName("Message").String()
	}

	return cErr
}
