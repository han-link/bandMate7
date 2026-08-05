package model

type Collection struct {
	BaseModel
	Titel        string        `json:"titel"`
	Performances []Performance `json:"performances" gorm:"many2many:collections_performances"`
}
