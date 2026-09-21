package main

import (
	"bandMate7/internal/service"
	"bandMate7/internal/store"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
)

func newTestApplication(t *testing.T, cfg config) *application {
	t.Helper()

	logger := zap.NewNop().Sugar()
	mockStore := store.NewMockStore()
	services := service.NewServices(&mockStore, cfg.baseURL.String())

	return &application{
		logger:  logger,
		store:   mockStore,
		config:  cfg,
		service: services,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}

func prepBody(t *testing.T, payload map[string]any) []byte {
	marshalled, err := json.Marshal(payload)

	if err != nil {
		t.Fatal(err)
	}

	return marshalled
}
