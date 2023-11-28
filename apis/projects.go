package apis

import (
	"net/http"

	"github.com/camarin24/go-studio/core"
	"github.com/camarin24/go-studio/services"
	"github.com/labstack/echo/v4"
)

type ProjectsApi struct {
	app     core.App
	service services.Service
}

func bindProjectsApi(app core.App, e *echo.Group) {
	api := &ProjectsApi{app, *services.NewService(app.DB(), app.Log())}
	group := e.Group("/projects")
	group.GET("", api.getProjects)
}

func (api *ProjectsApi) getProjects(c echo.Context) error {
	projects, err := api.service.GetProjects(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, projects)
}
