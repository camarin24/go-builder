package apis

import (
	"github.com/camarin24/go-studio/core"
	"github.com/camarin24/go-studio/ui"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const trailedAdminPath = "/_/"

func InitApp(app *core.App) (*echo.Echo, error) {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	bindUi(app, e)

	api := e.Group("/api")
	bindAdminApi(*app, api)
	bindProjectsApi(*app, api)

	return e, nil
}

func bindUi(app *core.App, e *echo.Echo) {
	e.GET(trailedAdminPath+"*",
		echo.StaticDirectoryHandler(ui.DistDirFS, false),
		uiCacheMiddleware(),
		middleware.Gzip())
}

func uiCacheMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().URL.Path != trailedAdminPath {
				c.Response().Header().Set("Cache-Control", "max-age=1209600, stale-while-revalidate=86400")
			}
			return next(c)
		}
	}
}
