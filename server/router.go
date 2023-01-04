package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"karintou8710/iNAZO-server/config"
	"karintou8710/iNAZO-server/controllers"
)

// NewRouter is constructor for router
func NewRouter() (*echo.Echo, error) {
	c := config.GetConfig()
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: c.GetStringSlice("app.cors"),
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	route := router.Group("/api")

	gradeDistributionController := controllers.NewGradeDistributionController()
	route.GET("/grade_distribution", gradeDistributionController.Search)

	return router, nil
}
