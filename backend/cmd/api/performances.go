package main

import (
	"bandMate7/internal/model"
	"bandMate7/internal/service"
	"bandMate7/internal/store"
	"errors"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type performanceKey string

const performanceCtx performanceKey = "performance"

// Delete performances
//
//	@Summary	Delete a performances
//	@Tags		performances
//	@Produce	json
//	@Param		id	path	string	true	"Performance ID"	Format(uuid)
//	@Success	204	"No Content"
//	@Failure	404	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/performances/{id} [delete]
func (app *application) deletePerformancesHandler(w http.ResponseWriter, r *http.Request) {
	performance := getPerformanceFromContext(r)
	ctx := r.Context()

	err := app.store.Performances.Delete(ctx, performance)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Get performances
//
//	@Summary	Get all performances
//	@Tags		performances
//	@Produce	json
//	@Param		desc	query		boolean	false	"Desc"
//	@Param		orderBy	query		string	false	"Order By"
//	@Success	200		{object}	[]model.Performance
//	@Failure	500		{object}	ErrorResponse
//	@Router		/performances [get]
func (app *application) getPerformancesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	performances, err := app.service.Performances.GetAll(ctx, r)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, performances); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Get resources
//
//	@Summary	Get all resources from a performance
//	@Tags		performances
//	@Produce	json
//	@Param		id	path		string	true	"Performance ID"	Format(uuid)
//	@Success	200	{object}	[]model.Resource
//	@Failure	404	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/performances/{id}/resources [get]
func (app *application) getResourcesByPerformanceHandler(w http.ResponseWriter, r *http.Request) {
	performance := getPerformanceFromContext(r)

	if err := app.jsonResponse(w, http.StatusOK, performance.Resources); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Get performance
//
//	@Summary	Get performance by id
//	@Tags		performances
//	@Produce	json
//	@Param		id	path		string	true	"Performance ID"	Format(uuid)
//	@Success	200	{object}	model.Performance
//	@Failure	404	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/performances/{id} [get]
func (app *application) getPerformanceHandler(w http.ResponseWriter, r *http.Request) {
	performance := getPerformanceFromContext(r)

	if err := app.jsonResponse(w, http.StatusOK, performance); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type CreatePerformancePayload struct {
	Name       string `json:"name" validate:"required,max=255"`
	Bpm        string `json:"bpm" validate:"omitempty,numeric,min=1"`
	UserRoleId string `json:"userRole" validate:"omitempty,uuid"`
}

func (payload CreatePerformancePayload) ParseBpm() (*int, error) {
	if payload.Bpm != "" {
		bpm, err := strconv.Atoi(payload.Bpm)
		if err != nil {
			return nil, err
		}
		return &bpm, nil
	}
	return nil, nil
}

func (payload CreatePerformancePayload) ParseUserRoleId() (*uuid.UUID, error) {
	if payload.UserRoleId != "" {
		userRoleId, err := uuid.Parse(payload.UserRoleId)
		if err != nil {
			return nil, err
		}
		return &userRoleId, nil
	}
	return nil, nil
}

// Create performance
//
//	@Summary	Create a new performance
//	@Tags		performances
//	@Accept		multipart/form-data
//	@Produce	json
//	@Param		name		formData	string	true	"Name"	minlength(1)	maxlength(255)
//	@Param		bpm			formData	int		false	"Bpm"
//	@Param		cover		formData	file	false	"Album/Performance cover"
//	@Param		file		formData	file	false	"Upload file"
//	@Param		userRoleId	formData	string	false	"Set user role"	Format(uuid)
//	@Success	201			{object}	model.Performance
//	@Failure	400			{object}	ErrorResponse
//	@Failure	500			{object}	ErrorResponse
//	@Router		/performances [post]
func (app *application) createPerformanceHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePerformancePayload
	payload.Name = r.FormValue("name")
	payload.Bpm = r.FormValue("bpm")
	payload.UserRoleId = r.FormValue("userRoleId")

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	bpm, err := payload.ParseBpm()
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	userRoleId, err := payload.ParseUserRoleId()
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	cover, coverHeader, err := r.FormFile("cover")
	if err != nil {
		if !errors.Is(err, http.ErrMissingFile) {
			app.internalServerError(w, r, err)
			return
		}
	}

	score, scoreHeader, err := r.FormFile("file")
	if err != nil {
		if !errors.Is(err, http.ErrMissingFile) {
			app.internalServerError(w, r, err)
			return
		}
	}

	ctx := r.Context()

	createdPerformance, err := app.service.Performances.Create(ctx, service.CreatePerformanceRequest{
		Name:        payload.Name,
		Bpm:         bpm,
		UserRoleId:  userRoleId,
		Cover:       cover,
		CoverHeader: coverHeader,
		Score:       score,
		ScoreHeader: scoreHeader,
	})
	if err != nil {
		switch {
		case errors.Is(err, service.ErrUserRoleIdRequiredForScore):
			app.badRequestResponse(w, r, err)
		case errors.Is(err, store.ErrUserRoleNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err = app.jsonResponse(w, http.StatusCreated, createdPerformance); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Create resource
//
//	@Summary	Create a new resource
//	@Tags		performances
//	@Accept		multipart/form-data
//	@Produce	json
//	@Param		id			path		string	true	"Performance ID"	Format(uuid)
//	@Param		role		formData	string	false	"Role"				Format(uuid)
//	@Param		resource	formData	file	false	"Bild"
//	@Success	201			{object}	model.Resource
//	@Failure	400			{object}	error
//	@Failure	404			{object}	error
//	@Router		/performances/{id}/resources [post]
func (app *application) createPerformanceResourceHandler(w http.ResponseWriter, r *http.Request) {
	roleParam := r.FormValue("role")

	roleId, err := uuid.Parse(roleParam)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	role, err := app.store.UserRoles.GetByID(ctx, roleId)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	file, header, err := r.FormFile("resource")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrMissingFile):
			app.badRequestResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}

	}

	performance := getPerformanceFromContext(r)

	err, resource := app.store.Resources.Create(ctx, file, header, performance, role)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, resource); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// TODO: Remove handler and replace with update performance handler
// Set cover
//
//	@Summary	Set a cover for a performance
//	@Tags		performances
//	@Accept		multipart/form-data
//	@Param		id		path		string	true	"Performance ID"	Format(uuid)
//	@Param		cover	formData	file	true	"Album/Performance cover"
//	@Success	201		"No Content"
//	@Failure	400		{object}	error
//	@Failure	404		{object}	error
//	@Router		/performances/{id}/cover [post]
func (app *application) setCoverHandler(w http.ResponseWriter, r *http.Request) {
	performance := getPerformanceFromContext(r)

	ctx := r.Context()

	file, header, err := r.FormFile("cover")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrMissingFile):
			if err = app.jsonResponse(w, http.StatusCreated, performance); err != nil {
				app.internalServerError(w, r, err)
				return
			}
			return
		default:
			app.internalServerError(w, r, err)
			return
		}

	}

	err, resource := app.store.Resources.Create(ctx, file, header, performance, nil)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err = app.store.Performances.SetCover(ctx, performance, resource); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, resource); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func getPerformanceFromContext(r *http.Request) *model.Performance {
	performance, _ := r.Context().Value(performanceCtx).(*model.Performance)
	return performance
}
