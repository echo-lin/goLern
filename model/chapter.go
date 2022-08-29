package model

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	Name string `gorm:"type:char(20);not null"`
	Pid  int    `gorm:"type:int(11);not null;"`
	Did  int    `gorm:"type:int(11);not null;"`
}
