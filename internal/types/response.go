package types

type Response[T any] struct {
	Data    T      `json:"data"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
} // @name Response

type Empty struct{}

type NoContent = Response[Empty] // @name NoContent
