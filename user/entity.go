package user

import "time"

type User struct {
	ID             uint      `gorm:"primary_key"`
	Name           string    `gorm:"column:name"`
	Job            string    `gorm:"column:job"`
	Email          string    `gorm:"column:email"`
	PasswordHash   string    `gorm:"column:password_hash"`
	AvatarFileName string    `gorm:"column:avatar_file_name"`
	Role           string    `gorm:"column:role"`
	CreatedAt      time.Time `gorm:"column:created_time;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_time;autoUpdateTime"`
}
