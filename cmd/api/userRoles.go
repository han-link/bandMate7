package main

import "net/http"

// Get user roles
//
//	@Summary	Get all user roles
//	@Tags		userRoles
//	@Success	200	{object}	[]model.UserRole
//	@Failure	400	{object}	error
//	@Failure	404	{object}	error
//	@Router		/userRoles [get]
func (app *application) getUserRolesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roles, err := app.store.UserRoles.GetAll(ctx)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err = app.jsonResponse(w, http.StatusOK, roles); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
