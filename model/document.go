package model

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	DocName string `gorm:"type:char(20);not null"`
}
