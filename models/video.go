package models

type Video struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Platform    string `gorm:"foreignKey:PlatformID"`
	Status      string `gorm:"foreignKey:StatusID"`
	Creator     string `gorm:"foreignKey:UserID"`
}

// TableName returns the table name for the Video model.
func (Video) TableName() string {
	return "videos"
}
