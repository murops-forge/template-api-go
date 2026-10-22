package api

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	h := NewHandler(slog.New(slog.NewTextHandler(io.Discard, nil)))

	tests := []struct {
		path    string
		status  int
		message string
	}{
		{"/healthz", http.StatusOK, ""},
		{"/hello", http.StatusOK, "hello, world"},
		{"/hello?name=forge", http.StatusOK, "hello, forge"},
		{"/hello?name=" + strings.Repeat("a", 65), http.StatusBadRequest, ""},
		{"/missing", http.StatusNotFound, ""},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			rec := httptest.NewRecorder()
			h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, tt.path, nil))

			if rec.Code != tt.status {
				t.Fatalf("status = %d, want %d", rec.Code, tt.status)
			}
			if tt.message == "" {
				return
			}
			var body map[string]string
			if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			if body["message"] != tt.message {
				t.Errorf("message = %q, want %q", body["message"], tt.message)
			}
		})
	}
}
