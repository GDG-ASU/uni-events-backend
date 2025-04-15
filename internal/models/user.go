package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    ClerkID string `gorm:"uniqueIndex" json:"clerk_id"`
    Name    string `json:"name"`
    Email   string `gorm:"uniqueIndex" json:"email"`
    Role    string `json:"role"`
}