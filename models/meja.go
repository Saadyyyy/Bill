package models

import "time"

type Meja struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Meja      string    `json:"meja" form:"meja"`
	CreatedAt time.Time `json:"created_at"`
}
