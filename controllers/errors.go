package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newErrorWithMsg(err error, msg string) error {
	msg += " (" + err.Error() + ")"
	return fmt.Errorf(msg)
}

// JSONErrorHandler returns errors in json format
func JSONErrorHandler(err error, c echo.Context) {
	var code int
	var msg string
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	} else {
		code = http.StatusInternalServerError
		msg = err.Error()
	}

	c.Logger().Error(err)

	if !c.Response().Committed {
		c.JSON(code, APIError{code, msg})
	}
}
