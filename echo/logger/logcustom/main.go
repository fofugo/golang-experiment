package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", Handler)
	e.Logger.SetLevel(0)
	t := time.Now()
	startTime := t.Format("2006_01_02__15_04")
	f, err := os.Create(startTime + ".log")
	if err != nil {
		return
	}
	e.Logger.SetOutput(f)
	e.Logger.SetHeader(`{"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}",` + `"file":"${short_file}","line":"${line}"}`)
	e.Start(":6080")
}
func Handler(c echo.Context) error {
	json := map[string]interface{}{
		"Hello": "world",
		"Name":  "John",
	}
	c.Logger().Info("fofugo")
	c.Logger().Infoj(json)
	return c.String(http.StatusOK, "success")
}
