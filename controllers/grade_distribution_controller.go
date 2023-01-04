package controllers

import (
	"karintou8710/iNAZO-server/models"
	"karintou8710/iNAZO-server/scope"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GradeDistributionController struct{}

func NewGradeDistributionController() *GradeDistributionController {
	return new(GradeDistributionController)
}

func (controller *GradeDistributionController) Search(c echo.Context) error {
	var gradeModel models.GradeDistribution
	var pagination scope.Pagination
	if err := c.Bind(&pagination); err != nil {
		return err
	}
	err := gradeModel.ListWithPagination(&pagination)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		pagination,
	))
}
