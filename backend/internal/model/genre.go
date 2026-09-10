package model

type Genre struct {
	BaseModel
	Titel        string        `json:"titel" gorm:"not null;unique"`
	Performances []Performance `json:"-" gorm:"many2many:genres_performances"`
}
