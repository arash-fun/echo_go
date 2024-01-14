package handlers

import (
	"net/http"

	"github.com/arashtest/cmd/api/service"
	"github.com/labstack/echo/v4"
)

func PostsIndexHandler(c echo.Context) error {

	data, err := service.GetData()
	if err != nil {
		c.String(http.StatusBadGateway, "unable to loadr")
	}
	res := make(map[string]any)
	res["status"] = "OK!"
	res["data"] = data
	return c.JSON(http.StatusOK, res)

}

func PostSigleHandler(c echo.Context) error {
	ip := c.Param("ip")
	ipx
}
