package store

import "net/http"

type FilteredResourcesQuery struct {
	Role string `json:"role"`
}

func (rq FilteredResourcesQuery) Parse(r *http.Request) (FilteredResourcesQuery, error) {
	qs := r.URL.Query()

	role := qs.Get("role")
	if role != "" {
		rq.Role = role
	}

	return rq, nil
}
