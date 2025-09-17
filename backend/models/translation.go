package models

import "time"

type Translation struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Key          string    `gorm:"type:varchar(255);not null;index:idx"`
	Text         string    `gorm:"type:text;not null" json:"text"`
	TranslatorID uint      `json:"translator_id"`
	LanguageID   uint      `json:"language_id"`
	ChapterID    uint      `json:"chapter_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
