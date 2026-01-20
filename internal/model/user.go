package model

type User struct {
	BaseModel
	Role UserRole `json:"role" gorm:"not null"`
}

type UserRole struct {
	BaseModel
	Name string `json:"name" gorm:"not null;unique"`
}
