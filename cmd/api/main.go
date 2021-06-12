package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/qashir-api/database/postgresql"

	urlH "github.com/ranggarifqi/qashir-api/src/url/handler"
	urlRepo "github.com/ranggarifqi/qashir-api/src/url/repository"
	urlUC "github.com/ranggarifqi/qashir-api/src/url/usecase"
)

func main() {
	e := echo.New()

	db := postgresql.InitDB()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	urlGroup := e.Group("/url")
	urlRepository := urlRepo.NewUrlRepository(db)
	urlUsecase := urlUC.NewUrlUsecase(urlRepository)
	urlH.NewUrlHandler(urlGroup, urlUsecase)

	e.Logger.Fatal(e.Start(":3000"))
}
