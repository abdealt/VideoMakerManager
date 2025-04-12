package models

import "time"

type Video struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	PlatformID  uint      `json:"platform_id" gorm:"not null"`
	Platform    Platform  `json:"platform,omitempty" gorm:"foreignKey:PlatformID"`
	StatusID    uint      `json:"status_id" gorm:"not null"`
	Status      Status    `json:"status,omitempty" gorm:"foreignKey:StatusID"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	Creator     User      `json:"creator,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName returns the table name for the Video model.
func (Video) TableName() string {
	return "videos"
}
