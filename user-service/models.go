package main
// gorm Model's definition

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	Email string
	Password string `sql:"type:text;"`
}