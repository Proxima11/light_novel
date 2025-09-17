package models

type Language struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Type string `gorm:"type:varchar(100);not null;unique" json:"language"`
}
