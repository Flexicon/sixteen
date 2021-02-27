package main

import (
	"net/http"

	"github.com/flexicon/sixteen/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = views.NewRenderer()

	e.GET("/", homeHandler)
	e.GET("/links", linksHandler)

	e.Debug = true
	e.Logger.Fatal(e.Start(":9000"))
}

// homeHandler - responds to the server root route
func homeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

// linksHandler - lists out helpful links
func linksHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "links.html", map[string]interface{}{
		"Links": []string{
			"https://golang.org/pkg/embed/",
			"https://echo.labstack.com/guide/templates",
			"https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/",
		},
	})
}
