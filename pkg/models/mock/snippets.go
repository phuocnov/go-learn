package mock

import (
	"time"

	"github.com/phuocnov/golang-webserver/pkg/models"
)

var mockSnippets = &models.Snippet{
	ID:      1,
	Title:   "First Snippet",
	Content: "This is the content of the first snippet.",
	Created: time.Now().Add(-48 * time.Hour),
	Expires: time.Now().Add(24 * time.Hour),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippets, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	snippets := []*models.Snippet{mockSnippets}
	return snippets, nil
}
