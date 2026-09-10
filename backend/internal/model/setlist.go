package model

import (
	"time"

	"github.com/google/uuid"
)

type Setlist struct {
	BaseModel
	Titel string               `json:"titel"`
	Items []SetlistPerformance `json:"items" gorm:"foreignKey:SetlistID"`
}

type SetlistPerformance struct {
	SetlistID     uuid.UUID `gorm:"primaryKey"`
	PerformanceID uuid.UUID `gorm:"primaryKey"`
	Position      int       `json:"position"`
	CreatedAt     time.Time
	Performance   Performance `json:"performance" gorm:"foreignKey:PerformanceID"`
}
