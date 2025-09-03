package cmd

import (
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/static", "static")

	app.GET("/write", func(ctx echo.Context) error {
		return echo.ErrBadRequest
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	app.Logger.Fatal(app.Start(port))
}
