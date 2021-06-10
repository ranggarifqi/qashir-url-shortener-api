package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/qashir-api/domain"
	"github.com/ranggarifqi/qashir-api/helper"
)

type urlHandler struct {
	urlUsecase domain.IUrlUsecase
}

func NewUrlHandler(g *echo.Group, urlUsecase domain.IUrlUsecase) {
	handler := &urlHandler{
		urlUsecase,
	}

	g.GET("/url/:id", handler.FindById)
}

func (h *urlHandler) FindById(c echo.Context) error {
	id := c.Param("id")
	res, err := h.urlUsecase.FindById(id)
	if err != nil {
		return helper.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, domain.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "Data fetched successfully",
		Data:       res,
	})
}
