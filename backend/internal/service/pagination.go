package service

import (
	"net/http"
	"strconv"
)

type PaginatedRequest struct {
	Desc    bool   `json:"desc"`
	OrderBy string `json:"orderBy" validate:"max=100"`
}

func (pr PaginatedRequest) Parse(r *http.Request) (PaginatedRequest, error) {
	qs := r.URL.Query()

	orderBy := qs.Get("orderBy")
	if orderBy != "" {
		pr.OrderBy = orderBy
	}

	desc := qs.Get("desc")
	if desc != "" {
		d, err := strconv.ParseBool(desc)
		if err != nil {
			return pr, nil
		}
		pr.Desc = d
	}

	return pr, nil
}
