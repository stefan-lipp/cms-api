package models

import (
	"errors"
	"time"
)

// Post represents a post in the CMS
type Post struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `gorm:"size:255;not null" json:"title" binding:"required"`
	Content   string    `gorm:"type:text;not null" json:"content" binding:"required"`
	Author    string    `gorm:"size:100" json:"author"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Media     []Media   `gorm:"many2many:post_media" json:"media"`
}

func (p *Post) Validate() error {
	if p.Title == "" {
		return errors.New("title is required")
	}

	if p.Content == "" {
		return errors.New("content is required")
	}

	return nil
}
