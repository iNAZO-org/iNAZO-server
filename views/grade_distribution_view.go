package views

import (
	"fmt"

	"karintou8710/iNAZO-server/models"
	"karintou8710/iNAZO-server/scope"
)

type GradeDistributionView struct {
	ID uint `json:"id"`

	Subject      string  `json:"subject"`
	SubTitle     string  `json:"subTitle"`
	Class        string  `json:"class"`
	Teacher      string  `json:"teacher"`
	Year         int     `json:"year"`
	Semester     int     `json:"semester"`
	Faculty      string  `json:"faculty"`
	StudentCount int     `json:"studentCount"`
	Gpa          float64 `json:"gpa"`

	ApCount int `json:"apCount"`
	ACount  int `json:"aCount"`
	AmCount int `json:"amCount"`
	BpCount int `json:"bpCount"`
	BCount  int `json:"bCount"`
	BmCount int `json:"bmCount"`
	CpCount int `json:"cpCount"`
	CCount  int `json:"cCount"`
	DCount  int `json:"dCount"`
	DmCount int `json:"dmCount"`
	FCount  int `json:"fCount"`
}

func NewGradeDistributionView(gradeDistibution *models.GradeDistribution) *GradeDistributionView {
	return &GradeDistributionView{
		ID: gradeDistibution.ID,

		Subject:      gradeDistibution.Subject,
		SubTitle:     gradeDistibution.SubTitle,
		Class:        gradeDistibution.Class,
		Teacher:      gradeDistibution.Teacher,
		Year:         gradeDistibution.Year,
		Semester:     gradeDistibution.Semester,
		Faculty:      gradeDistibution.Faculty,
		StudentCount: gradeDistibution.StudentCount,
		Gpa:          gradeDistibution.Gpa,

		ApCount: gradeDistibution.ApCount,
		ACount:  gradeDistibution.ACount,
		AmCount: gradeDistibution.AmCount,
		BpCount: gradeDistibution.BpCount,
		BCount:  gradeDistibution.BCount,
		BmCount: gradeDistibution.BmCount,
		CpCount: gradeDistibution.CpCount,
		CCount:  gradeDistibution.CCount,
		DCount:  gradeDistibution.DCount,
		DmCount: gradeDistibution.DmCount,
		FCount:  gradeDistibution.FCount,
	}
}

func GradeDistributionWithPaginationView(pagination *scope.Pagination) (*PaginationView, error) {
	gradeDistributionList, ok := pagination.Rows.([]*models.GradeDistribution)
	if !ok {
		return nil, fmt.Errorf("GradeDistributionWithPaginationView error: type check failed.")
	}

	var gradeDistributionViewList []*GradeDistributionView
	for _, gd := range gradeDistributionList {
		gradeDistributionViewList = append(gradeDistributionViewList, NewGradeDistributionView(gd))
	}

	return &PaginationView{
		Limit:      pagination.Limit,
		Page:       pagination.Page,
		TotalPages: pagination.TotalPages,
		TotalRows:  pagination.TotalRows,
		Rows:       gradeDistributionViewList,
	}, nil
}
