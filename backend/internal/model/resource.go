package model

import "github.com/google/uuid"

type Resource struct {
	BaseModel
	Type          string     `json:"type" gorm:"not null"`
	UserRoleID    *uuid.UUID `json:"-"`
	UserRole      *UserRole  `json:"user_role" extensions:"x-nullable"`
	Filename      string     `json:"filename"`
	SizeBytes     int64      `json:"-"`
	PerformanceID uuid.UUID  `json:"-" gorm:"index;not null"`
	Url           string     `json:"url" gorm:"-"`
} //	@name	Resource

func (r *Resource) SetUrl(baseUrl string) {
	if r == nil {
		return
	}
	r.Url = baseUrl + "/resources/" + r.ID.String()
}
