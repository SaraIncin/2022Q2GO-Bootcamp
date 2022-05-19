package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"myproject/first"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Root route => handler
	e.GET("/:id", func(c echo.Context) error {
		items := first.ReadCsv()
		var item = &first.Item{}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		for _, v := range items {
			if v.Id == id {
				item.Id = v.Id
				item.Treat = v.Treat
				item.Prewt = v.Prewt
				item.Postwt = v.Postwt

			}
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		return c.JSONPretty(http.StatusOK, item, "  ")

	})
	// Run Server
	e.Logger.Fatal(e.Start(":8000"))
}
