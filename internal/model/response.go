package model

import (
	"time"

	"github.com/chonlawit-odds/task-api/internal/constant"
)

type BaseResponse[DataType any] struct {
	Message string   `json:"message,omitempty"`
	Data    DataType `json:"data,omitempty"`
}

type BaseResponseList[DataType any] struct {
	Count   int      `json:"count"`
	Results DataType `json:"results"`
}

type ResponseItem struct {
	ID          uint                `json:"id"`
	Title       string              `json:"title"`
	Amount      float64             `json:"amount"`
	Quantity    uint                `json:"quantity"`
	Status      constant.ItemStatus `json:"status"`
	CreatedTime time.Time           `json:"createdTime"`
	UpdatedTime time.Time           `json:"updatedTime"`
}
