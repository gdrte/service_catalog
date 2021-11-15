package models

import (
	"gorm.io/gorm"
)

type Version struct {
	gorm.Model
	Ver       string
	ServiceID uint
}

type Service struct {
	gorm.Model
	Name        string
	Description string
	Versions    []Version `json:",omitempty"`
}
