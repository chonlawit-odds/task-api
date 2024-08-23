package model

import "github.com/chonlawit-odds/task-api/internal/constant"

type RequestCreateItem struct {
	Title    string  `validate:"required"`
	Amount   float64 `validate:"required,gt=0"`
	Quantity uint    `validate:"required,gt=0"`
}

type RequestFindItem struct {
	Statuses []constant.ItemStatus `form:"status[]" validate:"dive,oneof=PENDING APPROVED REJECTED"`
}

type RequestUpdateItem struct {
	Status constant.ItemStatus `validate:"oneof=PENDING APPROVED REJECTED"`
}
