package user

import "time"

// User defines the user model
type User struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"type:varchar(255)"`
	Occupation     string `gorm:"type:varchar(255)"`
	Email          string `gorm:"type:varchar(255);unique"`
	PasswordHash   string `gorm:"type:varchar(255)"` // Add PasswordHash field
	AvatarFileName string `gorm:"type:varchar(255)"`
	Role           string `gorm:"type:varchar(50)"`
	Token          string `gorm:"type:varchar(255)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
