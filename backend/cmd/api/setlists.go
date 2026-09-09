package main

import (
	"bandMate7/internal/model"
	"bandMate7/internal/service"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type setlistKey string

const setlistCtx setlistKey = "setlist"

func getSetlistFromContext(r *http.Request) *model.Setlist {
	setlist, _ := r.Context().Value(setlistCtx).(*model.Setlist)
	return setlist
}

type CreateSetlistPayload struct {
	Title          string      `json:"title" validate:"required"`
	PerformanceIds []uuid.UUID `json:"performanceIds"`
}

// Create setlist
//
//	@Summary	Creates a setlist
//	@Tags		setlists
//	@Accept		json
//	@Produce	json
//	@Param		payload	body		CreateSetlistPayload	true	"Setlist payload"
//	@Success	201		{object}	model.Setlist
//	@Failure	400		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/setlists [post]
func (app *application) createSetlistHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateSetlistPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	request := service.CreateSetlistRequest{
		Title:          payload.Title,
		PerformanceIds: payload.PerformanceIds,
	}

	setlist, err := app.service.Setlists.Create(ctx, request)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrPerformancesNotFound):
			app.badRequestResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusCreated, setlist); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Get setlists
//
//	@Summary	Get all setlists
//	@Tags		setlists
//	@Success	200	{object}	[]service.SelistWithoutPerformance
//	@Failure	500	{object}	ErrorResponse
//	@Router		/setlists [get]
func (app *application) getSetlistsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	setlists, err := app.service.Setlists.GetAll(ctx)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err = app.jsonResponse(w, http.StatusOK, setlists); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Get setlist
//
//	@Summary	Get setlist by id
//	@Tags		setlists
//	@Produce	json
//	@Param		id	path		string			true	"Setlists ID"	Format(uuid)
//	@Success	200	{object}	model.Setlist	"Setlist"
//	@Failure	404	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/setlists/{id} [get]
func (app *application) getSetlistHandler(w http.ResponseWriter, r *http.Request) {
	setlist := getSetlistFromContext(r)

	if err := app.jsonResponse(w, http.StatusOK, setlist); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Chane setlist order
//
//	@Summary	Chane setlist order
//	@Tags		setlists
//	@Produce	json
//	@Param		id		path		string						true	"Setlists ID"	Format(uuid)
//	@Param		payload	body		service.ChangeOrderPayload	true	"ChangeSetlistOrderPayload"
//	@Success	200		{object}	model.Setlist				"Setlist"
//	@Failure	400		{object}	ErrorResponse
//	@Failure	404		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/setlists/{id}/order [put]
func (app *application) changeSetlistOrderHandler(w http.ResponseWriter, r *http.Request) {
	setlist := getSetlistFromContext(r)
	var payload service.ChangeOrderPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	setlist, err := app.service.Setlists.ChangeOrder(ctx, payload, setlist)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrPerformancesNotFound):
			app.badRequestResponse(w, r, err)
			return
		case errors.Is(err, service.ErrInvalidInput):
			app.badRequestResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, setlist); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
