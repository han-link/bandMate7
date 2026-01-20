package main

import (
	"bandMate7/internal/store"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Get resource
//
//	@Summary	Get resource by id
//	@Tags		resources
//	@Param		id	path		string	true	"Resource ID"	Format(uuid)
//	@Failure	400	{object}	error
//	@Failure	404	{object}	error
//	@Router		/resources/{id} [get]
func (app *application) getResourceHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "resourceId")
	resourceId, err := uuid.Parse(idParam)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	resource, err := app.store.Resources.GetByID(ctx, resourceId)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, resource.Filename))
	http.ServeFile(w, r, "./resources/"+resource.Filename)
}
