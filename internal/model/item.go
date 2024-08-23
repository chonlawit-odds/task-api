package model

import (
	"time"

	"github.com/chonlawit-odds/task-api/internal/constant"
)

type Item struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Amount      float64
	Quantity    uint
	Status      constant.ItemStatus
	CreatedTime time.Time
	UpdatedTime time.Time
}
