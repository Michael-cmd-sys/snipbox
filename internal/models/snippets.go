package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	Title string
	Content string
	Created time.Time
	Expires time.Time
	ID int
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest(id int) ([]*Snippet, error) {
	return nil, nil
}