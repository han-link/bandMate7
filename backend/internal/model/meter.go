package model

type Meter struct {
	BaseModel
	Titel string `json:"titel" gorm:"not null;unique"`
}
