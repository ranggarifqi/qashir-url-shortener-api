package helper

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/qashir-api/domain"
)

// HandleError ...
func HandleError(c echo.Context, err error) error {

	code := http.StatusInternalServerError
	message := err.Error()

	if e, ok := err.(*echo.HTTPError); ok {
		code = e.Code
		message = fmt.Sprint(e.Message)
	}

	return c.JSON(code, domain.ErrorResponse{StatusCode: code, Message: message})
}
