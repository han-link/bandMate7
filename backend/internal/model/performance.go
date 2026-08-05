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

	ArtistID *uuid.UUID `json:"-"`
	Artist   *Artist    `json:"artist"`

	MeterID *uuid.UUID `json:"-"`
	Meter   *Meter     `json:"meter"`

	Resources   []Resource   `json:"resources" gorm:"foreignKey:PerformanceID;constraint:OnDelete:CASCADE;"`
	Setlists    []Setlist    `json:"-" gorm:"many2many:setlist_performances"`
	Collections []Collection `json:"-" gorm:"many2many:collections_performances"`
} //	@name	Performance
