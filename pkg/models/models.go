package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrDuplicateEmail     = errors.New("models: email already existed")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// SnippetModelInterface describes the behavior required by the application
// for working with snippets. Both the real MySQL SnippetModel and
// test mocks will implement this.
type SnippetModelInterface interface {
	Insert(string, string, string) (int, error)
	Get(id int) (*Snippet, error)
	Latest() ([]*Snippet, error)
}

// UserModelInterface describes the behavior required for user-related data.
type UserModelInterface interface {
	Insert(name, email, password string) error
	Authenticate(email, password string) (int, error)
	Get(id int) (*User, error)
}
