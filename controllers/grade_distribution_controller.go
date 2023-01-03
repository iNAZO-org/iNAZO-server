package controllers

import (
	"karintou8710/iNAZO-server/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GradeDistributionController struct{}

func NewGradeDistributionController() *GradeDistributionController {
	return new(GradeDistributionController)
}

func (controller *GradeDistributionController) Search(c echo.Context) error {
	var gradeModel models.GradeDistribution
	gradeDistributionList, err := gradeModel.All()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		gradeDistributionList,
	))
}
