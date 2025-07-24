package models

import "time"

type Platform struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Videos    []Video   `json:"videos,omitempty" gorm:"foreignKey:PlatformID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName returns the table name for the Platform model.
func (Platform) TableName() string {
	return "platforms"
}
