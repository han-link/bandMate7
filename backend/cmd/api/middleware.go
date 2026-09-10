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

func (app *application) resourceContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		resource.SetUrl(app.config.baseUrl)
		ctx = context.WithValue(ctx, resourceCtx, resource)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.logger.Debugw("Request", "method", r.Method, "path", r.URL.Path+"?"+r.URL.RawQuery)
		next.ServeHTTP(w, r)
	})
}

func (app *application) setlistContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "setlistId")
		setlistId, err := uuid.Parse(idParam)

		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		setlist, err := app.store.SetLists.GetByID(ctx, setlistId, store.WithSortedSetlist())
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}
		ctx = context.WithValue(ctx, setlistCtx, setlist)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
