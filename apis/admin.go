package apis

import (
	"net/http"

	"github.com/camarin24/go-studio/core"
	"github.com/camarin24/go-studio/core/database"
	"github.com/camarin24/go-studio/services"
	"github.com/camarin24/go-studio/types"
	"github.com/camarin24/go-studio/utils"
	"github.com/labstack/echo/v4"
)

type AdminApi struct {
	app     core.App
	service services.Service
}

func bindAdminApi(app core.App, e *echo.Group) {
	api := &AdminApi{app, *services.NewService(app.DB(), app.Log())}
	subGroup := e.Group("/admin")
	subGroup.GET("", api.testApi)
	subGroup.POST("/ping", api.ping)
}

func (api *AdminApi) testApi(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Hello": "World",
	})
}

func (api *AdminApi) ping(c echo.Context) error {
	var dbPing types.DbPing
	if err := c.Bind(&dbPing); err != nil {
		return utils.StandardError(c, err.Error())
	}
	fc := database.FactoryConfig{
		ConnStr:  dbPing.ConnStr,
		ConnType: dbPing.ConnType,
	}
	factory := fc.NewDbFactory()
	return c.JSON(200, echo.Map{"ping": factory.Ping()})
}
