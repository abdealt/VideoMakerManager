package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Password  string    `json:"password,omitempty" gorm:"not null"`
	Name      string    `json:"name"`
	Videos    []Video   `json:"videos,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName returns the table name for the User model.
func (User) TableName() string {
	return "users"
}
