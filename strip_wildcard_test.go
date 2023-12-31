package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestStripWildcardPrefix tests the StripWildcardPrefix function
func TestStripWildcardPrefix(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		prefix         string
		requestPath    string
		expectedPath   string
		expectedStatus int
	}{
		{
			name:           "Match static and dynamic segments",
			prefix:         "/s/{slug}",
			requestPath:    "/s/abc/test",
			expectedPath:   "/test",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Unmatched static segment",
			prefix:         "/s/{slug}",
			requestPath:    "/x/abc/test",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "No dynamic segment",
			prefix:         "/s/abc",
			requestPath:    "/s/abc/test",
			expectedPath:   "/test",
			expectedStatus: http.StatusOK,
		},
	}

	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != testCases[0].expectedPath {
			t.Errorf("Expected path '%s', got '%s'", testCases[0].expectedPath, r.URL.Path)
		}
	})

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest("GET", tc.requestPath, nil)
			rr := httptest.NewRecorder()

			handler := StripWildcardPrefix(tc.prefix, dummyHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
			}
		})
	}
}

func BenchmarkStripWildcardPrefix(b *testing.B) {
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	handler := StripWildcardPrefix("/s/{slug}", dummyHandler)

	req := httptest.NewRequest("GET", "/s/abc/test", nil)
	rr := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(rr, req)
	}
}

func BenchmarkHttpStripPrefix(b *testing.B) {
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	handler := http.StripPrefix("/s/abc", dummyHandler)

	req := httptest.NewRequest("GET", "/s/abc/test", nil)
	rr := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(rr, req)
	}
}
