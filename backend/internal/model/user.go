package model

type User struct {
	BaseModel
	Role UserRole `json:"role" gorm:"not null"`
} //	@name	User

type UserRole struct {
	BaseModel
	Titel string `json:"titel" gorm:"not null;unique"`
} //	@name	UserRole
