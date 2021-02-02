package models

import (
	"html"
	"strings"
	"time"
)

// Blog this the models define
type Blog struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Prepare this fuction is use while creating new entery in database
func (b *Blog) Prepare() {
	b.ID = 0
	b.Title = html.EscapeString(strings.TrimSpace(b.Title))
	b.Content = html.EscapeString(strings.TrimSpace(b.Content))
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}
