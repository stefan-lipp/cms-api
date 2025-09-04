package models

import "time"

// Media represents some media in the CMS
type Media struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	URL       string    `gorm:"size:255;not null" json:"url" binding:"required"`
	Type      string    `gorm:"size:50" json:"type" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
