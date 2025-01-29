package models

type Platform struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

// TableName returns the table name for the Platform model.
func (Platform) TableName() string {
	return "platforms"
}
