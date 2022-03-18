package main

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Phone string `gorm:"index"`
	Pass  string
	Token string
}
