package models

import "gorm.io/gorm"


type User struct {
    gorm.Model
    Nome  string `gorm:"size:255;not null" json:"nome"`
    Email string `gorm:"size:255;not null;uniqueIndex" json:"email"`
}
