package models

import "time"

type Novel struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null;unique" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	AuthorID    uint      `json:"author_id"`
	CoverImage  string    `gorm:"type:varchar(255)" json:"cover_image"`
	Genre       string    `gorm:"type:varchar(100)" json:"genre"`
	Status      string    `gorm:"type:varchar(50)" json:"status"` // ongoing, completed, hiatus, cancelled
	Views       uint      `json:"views"`
	Likes       uint      `json:"likes"`
	Favorites   uint      `json:"favorites"`
	Chapters    []Chapter `gorm:"foreignKey:NovelID" json:"chapters"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
