package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// TableName returns the table name for the User model.
func (User) TableName() string {
	return "users"
}
