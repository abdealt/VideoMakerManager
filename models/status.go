package models

type Status struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

// TableName returns the table name for the State model.
func (Status) TableName() string {
	return "status"
}
