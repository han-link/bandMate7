package main

import (
	"bandMate7/internal/model"
	"net/http"
)

// Get artists
//
//	@Summary	Get all artists
//	@Tags		artists
//	@Success	200	{object}	[]model.Artist
//	@Failure	500	{object}	ErrorResponse
//	@Router		/artists [get]
func (app *application) getArtistsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roles, err := app.store.Artists.GetAll(ctx)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err = app.jsonResponse(w, http.StatusOK, roles); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type CreateArtistPayload struct {
	Title string `json:"title" validate:"required"`
}

// Create artist
//
//	@Summary	Creates a artist
//	@Tags		artists
//	@Accept		json
//	@Produce	json
//	@Param		payload	body		CreateArtistPayload	true	"Artist payload"
//	@Success	201		{object}	model.Artist
//	@Failure	400		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/artists [post]
func (app *application) createArtistHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateArtistPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	artist := &model.Artist{
		Titel: payload.Title,
	}

	ctx := r.Context()

	if err := app.store.Artists.Create(ctx, artist); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, artist); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
