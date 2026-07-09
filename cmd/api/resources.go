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
//	@Produce	octet-stream
//	@Param		id	path		string	true	"Resource ID"	Format(uuid)
//	@Success	200	{string}	binary	"Resource file"
//	@Failure	404	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
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

	object, err := app.store.Resources.Open(ctx, resource)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}
	defer object.Close()

	contentType := object.ContentType
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%q", resource.Filename))
	http.ServeContent(w, r, resource.Filename, object.LastModified, object)
}
