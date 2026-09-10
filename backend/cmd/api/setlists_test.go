package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetSetlist(t *testing.T) {
	cfg := config{}
	app := newTestApplication(t, cfg)
	mux := app.mount()

	t.Run("should return setlist", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/setlists/eed3fbc0-6d6d-4e14-9a0f-17741ccdd1c5", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusOK, rr.Code)
	})

	t.Run("should return all setlists", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/setlists", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusOK, rr.Code)
	})

	t.Run("should return error on malformed uuid", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/setlists/eed3fbc0-6d6d", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("create a setlist", func(t *testing.T) {
		payload := map[string]any{
			"title": "Test 123",
		}

		req, err := http.NewRequest(http.MethodPost, "/api/v1/setlists", bytes.NewReader(prepBody(t, payload)))
		if err != nil {
			t.Fatal(err)
		}
		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusCreated, rr.Code)
	})

	t.Run("setlist create should fail on missing title", func(t *testing.T) {
		payload := map[string]any{
			"performanceIds": []string{},
		}

		req, err := http.NewRequest(http.MethodPost, "/api/v1/setlists", bytes.NewReader(prepBody(t, payload)))
		if err != nil {
			t.Fatal(err)
		}
		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("setlist create should fail on null title", func(t *testing.T) {
		payload := map[string]any{
			"title":          nil,
			"performanceIds": []string{},
		}

		req, err := http.NewRequest(http.MethodPost, "/api/v1/setlists", bytes.NewReader(prepBody(t, payload)))
		if err != nil {
			t.Fatal(err)
		}
		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("setlist create should at least include on character", func(t *testing.T) {
		payload := map[string]any{
			"title":          "",
			"performanceIds": []string{},
		}

		req, err := http.NewRequest(http.MethodPost, "/api/v1/setlists", bytes.NewReader(prepBody(t, payload)))
		if err != nil {
			t.Fatal(err)
		}
		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("setlist create should at least include on character", func(t *testing.T) {
		payload := []byte(`{{`)

		req, err := http.NewRequest(http.MethodPost, "/api/v1/setlists", bytes.NewReader(payload))
		if err != nil {
			t.Fatal(err)
		}
		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})
}
