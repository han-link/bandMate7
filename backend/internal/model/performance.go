package model

import "github.com/google/uuid"

type Performance struct {
	BaseModel
	Titel    string `json:"titel"`
	Bpm      *int   `json:"bpm"`
	Released *int   `json:"released"`
	Duration *int   `json:"duration"`

	CoverID *uuid.UUID `json:"-"`
	Cover   *Resource  `json:"cover" gorm:"foreignKey:CoverID;references:ID;-:migration"`

	Resources []Resource `json:"resources" gorm:"foreignKey:PerformanceID;constraint:OnDelete:CASCADE;"`
	Setlists  []Setlist  `json:"-" gorm:"many2many:setlist_performances"`
} //	@name	Performance
