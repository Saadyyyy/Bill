package models

import "time"

type Menu struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Item      string    `json:"item" form:"item"`
	Harga     float64   `json:"harga" form:"harga"`
	CreatedAt time.Time `json:"created_at"`
}
