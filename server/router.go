package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"karintou8710/iNAZO-server/controllers"
)

// NewRouter is constructor for router
func NewRouter() (*echo.Echo, error) {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	route := router.Group("/api")

	gradeDistributionController := controllers.NewGradeDistributionController()
	route.GET("/grade_distribution", gradeDistributionController.Search)

	return router, nil
}
