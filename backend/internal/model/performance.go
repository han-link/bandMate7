package model

import "github.com/google/uuid"

type Performance struct {
	BaseModel
	Titel    string `json:"titel"`
	Bpm      *int   `json:"bpm" extensions:"x-nullable"`
	Released *int   `json:"released" extensions:"nullable"`
	Duration *int   `json:"duration" extensions:"x-nullable"`

	CoverID *uuid.UUID `json:"-"`
	Cover   *Resource  `json:"cover" gorm:"foreignKey:CoverID;references:ID;-:migration" extensions:"x-nullable"`

	ArtistID *uuid.UUID `json:"-"`
	Artist   *Artist    `json:"artist" extensions:"x-nullable"`

	MeterID *uuid.UUID `json:"-"`
	Meter   *Meter     `json:"meter" extensions:"x-nullable"`

	Resources    []Resource           `json:"resources" gorm:"foreignKey:PerformanceID;constraint:OnDelete:CASCADE;"`
	SetlistItems []SetlistPerformance `json:"-" gorm:"foreignKey:PerformanceID"`
	Collections  []Collection         `json:"-" gorm:"many2many:collections_performances"`
	Genres       []Genre              `json:"genres" gorm:"many2many:genres_performances"`
} //	@name	Performance

func (p *Performance) SetUrls(baseUrl string) {
	if p == nil {
		return
	}
	if p.Cover != nil {
		p.Cover.SetUrl(baseUrl)
	}
	for i := range p.Resources {
		p.Resources[i].SetUrl(baseUrl)
	}
}
