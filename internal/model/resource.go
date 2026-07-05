package model

import "github.com/google/uuid"

type ResourceType string

const (
	ResourceTypeImage    ResourceType = "image"
	ResourceTypeVideo    ResourceType = "video"
	ResourceTypeAudio    ResourceType = "audio"
	ResourceTypeDocument ResourceType = "document"
)

type Resource struct {
	BaseModel
	Type          ResourceType `json:"type" gorm:"type:resource_type;not null"`
	UserRoleID    *uuid.UUID   `json:"-"`
	UserRole      *UserRole    `json:"role"`
	Filename      string       `json:"filename"`
	SizeBytes     int64        `json:"-"`
	PerformanceID uuid.UUID    `json:"-" gorm:"index;not null"`
	Url           string       `json:"url" gorm:"-"`
} //	@name	Resource

func (r *Resource) SetUrl(baseUrl string) {
	r.Url = baseUrl + "/api/v1/resources/" + r.ID.String()
}
