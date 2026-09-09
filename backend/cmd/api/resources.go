package main

import (
	"bandMate7/internal/model"
	"bandMate7/internal/store"
	"errors"
	"fmt"
	"net/http"
)

type resourceKey string

const resourceCtx resourceKey = "resource"

func getResourceFromContext(r *http.Request) *model.Resource {
	resource, _ := r.Context().Value(resourceCtx).(*model.Resource)
	return resource
}

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
	resource := getResourceFromContext(r)

	ctx := r.Context()
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

	defer func() {
		if err := object.Close(); err != nil {
			app.logger.Fatal(err)
		}
	}()

	contentType := object.ContentType
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%q", resource.Filename))
	http.ServeContent(w, r, resource.Filename, object.LastModified, object)
}

// Get resource meta
//
//	@Summary	Get resource by id
//	@Tags		resources
//	@Produce	json
//	@Param		id	path		string			true	"Resource ID"	Format(uuid)
//	@Success	200	{object}	model.Resource	"Resource Meta"
//	@Failure	404	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/resources/{id}/meta [get]
func (app *application) getResourceMetaHandler(w http.ResponseWriter, r *http.Request) {
	resource := getResourceFromContext(r)

	if err := app.jsonResponse(w, http.StatusOK, resource); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
