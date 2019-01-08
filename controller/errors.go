package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newErrorWithMsg(err error, msg string) error {
	msg += " (" + err.Error() + ")"
	return fmt.Errorf(msg)
}

// error handler which returns errors in json format
func JSONErrorHandler(err error, c echo.Context) {
	var code int
	var msg string
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
		// msg = string(he.Message)
	} else {
		code = http.StatusInternalServerError
		msg = err.Error()
	}

	apierr := APIError{
		code,
		msg,
	}

	c.Logger().Error(err)

	if !c.Response().Committed {
		c.JSON(code, apierr)
	}
}
