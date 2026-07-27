package model

import "github.com/google/uuid"

type Performance struct {
	BaseModel
	Name string `json:"name"`
	Bpm  *int   `json:"bpm"`

	CoverID *uuid.UUID `json:"-"`
	Cover   *Resource  `json:"cover" gorm:"foreignKey:CoverID;references:ID;-:migration"`

	Resources []Resource `json:"resources" gorm:"foreignKey:PerformanceID;constraint:OnDelete:CASCADE;"`
} //	@name	Performance
