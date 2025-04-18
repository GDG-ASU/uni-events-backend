package models

import "gorm.io/gorm"

type Club struct {
	gorm.Model
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Owners      []ClubOwner  `gorm:"foreignKey:ClubID" json:"owners"`
}
