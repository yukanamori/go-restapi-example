package entity

import (
	"time"

	"gorm.io/gorm"
)

// User はユーザです。
type User struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	Username     string         `gorm:"unique;not null;varchar(50)" json:"username"`
	Password     string         `gorm:"not null;varchar(255)" json:"-"`
	Email        string         `gorm:"unique;not null;varchar(255)" json:"email"`
	FirstName    string         `gorm:"not null;varchar(50)" json:"first_name"`
	LastName     string         `gorm:"not null;varchar(50)" json:"last_name"`
	ProfileImage string         `gorm:"text" json:"profile_image"`
	Version      uint           `gorm:"default:0" json:"version"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
