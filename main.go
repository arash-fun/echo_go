package main

import (
	"github.com/labstack/echo/v4"
)

type DNS struct {
	ip     string
	name   int
	status string
}

func (b *DNS) getter(value string) *DNS {
	b.name = value
	return b
}

func main() {
	e := echo.New()
	e.GET("/health-check", handlers.healthCheckHandler)
	e.GET("/posts", handlers.postsIndexHandler)
	e.GET("post/:id", handlers.postSigleHandler)

	e.Logger.Fatal(e.Start(":5004"))
}
