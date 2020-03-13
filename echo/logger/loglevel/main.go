package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", Handler)
	e.Logger.SetLevel(3)
	e.Start(":6080")
}
func Handler(c echo.Context) error {
	c.Logger().Debug("This is Debug")
	c.Logger().Info("This is Info")
	c.Logger().Warn("This is Warn")
	c.Logger().Error("This is Error")
	c.Logger().Panic("This is Panic")
	c.Logger().Fatal("This is Fatal")
	return c.String(http.StatusOK, "success")
}
