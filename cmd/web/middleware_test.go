package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()

	frameOptions := rs.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
		t.Errorf("want X-Frame-Options to equal 'deny'; got %q", frameOptions)
	}

	xssProtection := rs.Header.Get("X-XSS-Protection")
	if xssProtection != "1; mainmode=block" {
		t.Errorf("want X-XSS-Protection to equal '1; mainmode=block'; got %q", xssProtection)
	}

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal 'OK'")
	}
}
