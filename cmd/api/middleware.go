package main

import (
	"bandMate7/internal/store"
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (app *application) performanceContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "performanceId")
		performanceId, err := uuid.Parse(idParam)

		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		performance, err := app.store.Performances.GetByID(ctx, performanceId)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}
		ctx = context.WithValue(ctx, performanceCtx, performance)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
