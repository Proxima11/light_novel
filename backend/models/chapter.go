package models

import "time"

type Chapter struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	NovelID   uint      `json:"novel_id"`
	Order     uint      `json:"order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
