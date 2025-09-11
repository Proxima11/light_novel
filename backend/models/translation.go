package models

type Translation struct {
	ID  uint   `gorm:"primaryKey" json:"id"`
	Key string `gorm:"type:varchar(255);not null;index:idx"`
}
