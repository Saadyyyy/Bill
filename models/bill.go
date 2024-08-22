package models

import "time"

type Bill struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" form:"name"`
	Amount    float64   `json:"amount" form:"amount"`
	Hours     float64   `json:"hours" form:"hours"`
	MejaId    uint      `json:"mejaId" form:"mejaId"`
	MenuId    uint      `json:"menuId" form:"menuId"`
	CreatedAt time.Time `json:"created_at"`
	Menu      Menu      `json:"menu" gorm:"foreignKey:MenuId"`
	Meja      Meja      `json:"meja" gorm:"foreignKey:MejaId"`
}
