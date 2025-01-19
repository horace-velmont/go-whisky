// Package http provides HTTP-specific error handling functionality
package http

import (
	"errors"
	"net/http"

	"github.com/samber/lo"
)

const errorNameNotFound = "notFoundError"

func NotFound(errs ...error) error {
	return newResponseError(errorNameNotFound, http.StatusNotFound, errs...)
}

func IsNotFound(err error) bool {
	var re *ResponseError
	if errors.As(err, &re) {
		return re.statusCode == http.StatusNotFound &&
			re.errorName == errorNameNotFound
	}
	return false
}

type ResponseError struct {
	errs       []error
	errorName  string
	statusCode int
}

func (e *ResponseError) Error() string {
	return e.errorName
}

func newResponseError(errorName string, statusCode int, errs ...error) *ResponseError {
	return &ResponseError{
		lo.Filter(
			errs, func(err error, _ int) bool {
				return err != nil
			},
		),
		errorName,
		statusCode,
	}
} 
