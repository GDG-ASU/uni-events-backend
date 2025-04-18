package models

import "gorm.io/gorm"

type ClubOwner struct {
	gorm.Model
	UserID  uint `gorm:"uniqueIndex:idx_club_user"`
	User   User `gorm:"foreignKey:UserID" json:"user"`

	ClubID  uint `gorm:"uniqueIndex:idx_club_user"`
	Club   Club `gorm:"foreignKey:ClubID" json:"club"`
}
