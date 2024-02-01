package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model

	Name       string `valid:"required~Name is required"`
	URL        string `valid:"required~URL is required, url~Url URL is invalid"`
	EmployeeID string `valid:"required~Name is required, matches(^[EM]{2}\\d{10}$)"`
}
