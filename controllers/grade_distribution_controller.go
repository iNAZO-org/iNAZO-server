package controllers

import (
	"karintou8710/iNAZO-server/models"
	"karintou8710/iNAZO-server/scope"
	"karintou8710/iNAZO-server/views"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GradeDistributionController struct{}

type RequestQuery struct {
	scope.PaginationQuery
	Sort   string `query:"sort"`
	Search string `query:"search"`
}

func NewGradeDistributionController() *GradeDistributionController {
	return new(GradeDistributionController)
}

func (r *RequestQuery) NewPaginationFromRequestQuery() *scope.Pagination {
	return &scope.Pagination{
		Limit: r.Limit,
		Page:  r.Page,
	}
}

func (controller *GradeDistributionController) Search(c echo.Context) error {
	var gradeModel models.GradeDistribution
	var requestQuery RequestQuery
	if err := c.Bind(&requestQuery); err != nil {
		return err
	}
	pagination := requestQuery.NewPaginationFromRequestQuery()
	err := gradeModel.ListWithPagination(pagination, requestQuery.Search, requestQuery.Sort)
	if err != nil {
		return err
	}

	result, err := views.GradeDistributionWithPaginationView(pagination)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		result,
	))
}
