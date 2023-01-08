package models

import (
	"karintou8710/iNAZO-server/database"
	"karintou8710/iNAZO-server/scope"

	"gorm.io/gorm"
)

type GradeDistribution struct {
	gorm.Model

	Subject      string  `gorm:"uniqueIndex:unique_column"`
	SubTitle     string  `gorm:"uniqueIndex:unique_column"`
	Class        string  `gorm:"uniqueIndex:unique_column"`
	Teacher      string  `gorm:"uniqueIndex:unique_column"`
	Year         int     `gorm:"uniqueIndex:unique_column"`
	Semester     int     `gorm:"uniqueIndex:unique_column"`
	Faculty      string  `gorm:"uniqueIndex:unique_column"`
	StudentCount int     `gorm:"uniqueIndex:unique_column"`
	Gpa          float64 `gorm:"uniqueIndex:unique_column"`

	ApCount int // A+の人数
	ACount  int // A
	AmCount int // A-
	BpCount int // B+
	BCount  int // B
	BmCount int // B-
	CpCount int // C+
	CCount  int // C
	DCount  int // D
	DmCount int // D-
	FCount  int // F
}

func SortScope(sortQuery string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if sortQuery == "gpa" {
			db = db.Order("gpa ASC")
		} else if sortQuery == "-gpa" {
			db = db.Order("gpa DESC")
		} else if sortQuery == "failure" {
			db = db.Order("(d_count + dm_count + f_count) * 100 / student_count ASC")
		} else if sortQuery == "-failure" {
			db = db.Order("(d_count + dm_count + f_count) * 100 / student_count DESC")
		} else if sortQuery == "a_band" {
			db = db.Order("(ap_count + a_count + am_count) * 100 / student_count ASC")
		} else if sortQuery == "-a_band" {
			db = db.Order("(ap_count + a_count + am_count) * 100 / student_count DESC")
		} else if sortQuery == "-f" {
			db = db.Order("f_count * 100 / student_count DESC")
		}
		return db.Order("year DESC, semester DESC")
	}
}

func SearchScope(searchQuery string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if searchQuery == "" {
			return db
		}

		return db.Where("translate_case(subject) LIKE '%' || translate_case(?) || '%'", searchQuery).
			Or("translate_case(sub_title) LIKE '%' || translate_case(?) || '%'", searchQuery).
			Or("translate_case(class) LIKE '%' || translate_case(?) || '%'", searchQuery).
			Or("translate_case(teacher) LIKE '%' || translate_case(?) || '%'", searchQuery).
			Or("translate_case(year::TEXT) LIKE '%' || translate_case(?) || '%'", searchQuery).
			Or("translate_case(semester::TEXT) LIKE '%' || translate_case(?) || '%'", searchQuery).
			Or("translate_case(faculty) LIKE '%' || translate_case(?) || '%'", searchQuery)
	}
}

func (model *GradeDistribution) ListWithPagination(pagination *scope.Pagination, searchQuery, sortQuery string) error {
	var gradeDitributionList []*GradeDistribution
	db := database.GetDB().Model(gradeDitributionList)

	err := db.Scopes(
		SortScope(sortQuery),
		SearchScope(searchQuery),
		scope.PaginateScope(pagination),
	).
		Find(&gradeDitributionList).Error
	pagination.Rows = gradeDitributionList
	return err
}
