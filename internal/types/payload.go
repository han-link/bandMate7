package types

type CreatePerformancePayload struct {
	Name string `json:"name" validate:"required,max=255"`
	Bpm  int    `json:"bpm"`
} // @name CreatePerformancePayload
