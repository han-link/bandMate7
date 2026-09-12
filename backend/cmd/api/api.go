package main

import (
	"bandMate7/internal/docs"
	"bandMate7/internal/service"
	"bandMate7/internal/store"
	"fmt"
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
	host        string
	baseUrl     string
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
	r.Use(app.requestLogger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
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
				r.Use(app.resourceContextMiddleware)
				r.Get("/", app.getResourceHandler)
				r.Get("/meta", app.getResourceMetaHandler)
			})
		})
		r.Route("/userRoles", func(r chi.Router) {
			r.Get("/", app.getUserRolesHandler)
		})
		r.Route("/artists", func(r chi.Router) {
			r.Get("/", app.getArtistsHandler)
			r.Post("/", app.createArtistHandler)
		})
		r.Route("/setlists", func(r chi.Router) {
			r.Post("/", app.createSetlistHandler)
			r.Get("/", app.getSetlistsHandler)
			r.Route("/{setlistId}", func(r chi.Router) {
				r.Use(app.setlistContextMiddleware)
				r.Get("/", app.getSetlistHandler)
				r.Put("/order", app.changeSetlistOrderHandler)
			})
		})
	})
	return r
}

func (app *application) run(mux http.Handler) error {
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.host + app.config.addr
	docs.SwaggerInfo.BasePath = "/api/v1"

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app.logger.Infow(fmt.Sprintf("Server has started at %s%s", app.config.host, app.config.addr))
	app.logger.Infow(fmt.Sprintf("Docs are available at %s%s%s", app.config.host, app.config.addr, "/api/v1/swagger/"))

	return srv.ListenAndServe()
}
