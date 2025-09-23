package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal 'OK'; got %q", string(body))
	}
}

func TestShowSnippet(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name       string
		urlPath    string
		wantStatus int
		wantBody   []byte
	}{
		{"Valid ID", "/snippet/1", http.StatusOK, []byte("This is the content of the first snippet.")},
		{"Non-existent ID", "/snippet/2", http.StatusNotFound, nil},
		{"Negative ID", "/snippet/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/snippet/1.5", http.StatusNotFound, nil},
		{"String ID", "/snippet/abc", http.StatusNotFound, nil},
		{"Empty ID", "/snippet/", http.StatusNotFound, nil},
		{"Trailing slash", "/snippet/1/", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)
			if code != tt.wantStatus {
				t.Errorf("want %d; got %d", tt.wantStatus, code)
			}
			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body to contain %q; got %q", tt.wantBody, body)
			}
		})
	}
}

// func TestSignupUSer(t *testing.T) {
// 	app := newTestApplication(t)
// 	ts := newTestServer(t, app.routes())
// 	defer ts.Close()
//
// 	_, _, body := ts.get(t, "/user/signup")
// 	csrfTokenRX := extractCSRFToken(t, body)
// 	t.Log(csrfTokenRX)
//
// 	tests := []struct {
// 		name         string
// 		userName     string
// 		userEmail    string
// 		userPassword string
// 		csrfToken    string
// 		wantStatus   int
// 		wantBody     []byte
// 	}{
// 		{"Valid submission", "Bob", "bob@example.com", "validPa$$word", csrfTokenRX, http.StatusSeeOther, nil},
// 		{"Empty name", "", "empty@example.com", "validPa$$word", csrfTokenRX, http.StatusOK, []byte("This field cannot be blank")},
// 		{"Empty email", "Empty", "", "validPa$$word", csrfTokenRX, http.StatusOK, []byte("This field cannot be blank")},
// 		{"Empty password", "Empty", "empty@example.com", "", csrfTokenRX, http.StatusOK, []byte("This field cannot be blank")},
// 		{"Invalid email (missing @)", "Invalid", "invalidexample.com", "validPa$$word", csrfTokenRX, http.StatusOK, []byte("This field is not a valid email address")},
// 		{"Invalid email (missing domain)", "Invalid", "invalid@", "validPa$$word", csrfTokenRX, http.StatusOK, []byte("This field is not a valid email address")},
// 		{"Short password", "Short", "short@gmail.com", "short", csrfTokenRX, http.StatusOK, []byte("This field is too short (minimum is 10 characters)")},
// 		{"Duplicate email", "Duplicate", "dupe@example.com", "validPa$$word", csrfTokenRX, http.StatusOK, []byte("Email address is already in use")},
// 		{"Invalid CSRF token", "CSRF", "", "validPa$$word", "", http.StatusBadRequest, nil},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			form := url.Values{}
// 			form.Add("name", tt.userName)
// 			form.Add("email", tt.userEmail)
// 			form.Add("password", tt.userPassword)
// 			form.Add("csrf_token", tt.csrfToken)
//
// 			code, _, body := ts.postForm(t, "/user/signup", form)
// 			if code != tt.wantStatus {
// 				t.Errorf("want %d; got %d", tt.wantStatus, code)
// 			}
// 			if !bytes.Contains(body, tt.wantBody) {
// 				t.Errorf("want body to contain %q; got %q", tt.wantBody, body)
// 			}
// 		})
// 	}
// }
