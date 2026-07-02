package main

import (
	"bandMate7/internal/docs"
	"bandMate7/internal/service"
	"bandMate7/internal/store"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"go.uber.org/zap"
)

type application struct {
	config  config
	logger  *zap.SugaredLogger
	store   store.Storage
	service service.Services
}

type config struct {
	addr        string
	apiURL      string
	resourceDir string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}

	r.Use(cors.Handler(corsOptions))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("doc.json")))
		r.Route("/performances", func(r chi.Router) {
			r.Post("/", app.createPerformanceHandler)
			r.Get("/", app.getPerformancesHandler)
			r.Route("/{performanceId}", func(r chi.Router) {
				r.Use(app.performanceContextMiddleware)
				r.Delete("/", app.deletePerformancesHandler)
				r.Get("/", app.getPerformanceHandler)
				r.Post("/cover", app.setCoverHandler)
				r.Route("/resources", func(r chi.Router) {
					r.Post("/", app.createPerformanceResourceHandler)
					r.Get("/", app.getResourcesByPerformanceHandler)
				})
			})
		})
		r.Route("/resources", func(r chi.Router) {
			r.Route("/{resourceId}", func(r chi.Router) {
				r.Get("/", app.getResourceHandler)
			})
		})
		r.Route("/userRoles", func(r chi.Router) {
			r.Get("/", app.getUserRolesHandler)
		})
	})
	return r
}

func (app *application) run(mux http.Handler) error {
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL + app.config.addr
	docs.SwaggerInfo.BasePath = "/api/v1"

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app.logger.Infow("Server has started at", "addr", app.config.addr)

	return srv.ListenAndServe()
}
