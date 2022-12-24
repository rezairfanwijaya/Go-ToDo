package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type responseAPISuccess struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type responseAPIFailed struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

func ResponseAPI(status string, data, message interface{}, IsError bool) interface{} {
	if IsError {
		return responseAPIFailed{
			Status:  status,
			Message: message,
		}
	}

	return responseAPISuccess{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func ErrorBinding(err error) []string {
	var myerr []string

	for _, e := range err.(validator.ValidationErrors) {
		errMessage := fmt.Sprintf("error on filed: %v, condition: %v", e.Field(), e.ActualTag())
		myerr = append(myerr, errMessage)
	}

	return myerr
}

func ValidateID(id int) error {
	if id <= 0 {
		errMsg := fmt.Sprintf("id must grather then 0")
		return errors.New(errMsg)
	}

	return nil
}
