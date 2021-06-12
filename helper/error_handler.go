package helper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/qashir-api/src/response"
)

func HandleError(msg string, err error) {
	if err != nil {
		log.Fatalf("%v: %v\n", msg, err)
	}
}

// HandleHttpError ...
func HandleHttpError(c echo.Context, err error) error {

	code := http.StatusInternalServerError
	message := err.Error()

	if e, ok := err.(*echo.HTTPError); ok {
		code = e.Code
		message = fmt.Sprint(e.Message)
	}

	return c.JSON(code, response.ErrorResponse{StatusCode: code, Message: message})
}
