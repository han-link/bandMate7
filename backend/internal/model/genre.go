package model

type Genre struct {
	BaseModel
	Titel        string        `json:"titel"`
	Performances []Performance `json:"-" gorm:"many2many:genres_performances"`
}
