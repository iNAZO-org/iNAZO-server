package models

import (
	"karintou8710/iNAZO-server/database"

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

func (model *GradeDistribution) All() ([]*GradeDistribution, error) {
	var gradeDitributionList []*GradeDistribution
	db := database.GetDB()
	err := db.Find(&gradeDitributionList).Error
	return gradeDitributionList, err
}
