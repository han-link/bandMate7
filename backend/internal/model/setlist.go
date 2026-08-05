package model

import (
	"time"

	"github.com/google/uuid"
)

type Setlist struct {
	BaseModel
	Titel        string        `json:"titel"`
	Performances []Performance `gorm:"many2many:setlist_performances"`
}

type SetlistPerformance struct {
	SetlistID     uuid.UUID `gorm:"primaryKey"`
	PerformanceID uuid.UUID `gorm:"primaryKey"`
	Position      int
	CreatedAt     time.Time
}
