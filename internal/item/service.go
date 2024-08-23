package item

import (
	"time"

	"github.com/chonlawit-odds/task-api/internal/constant"
	"github.com/chonlawit-odds/task-api/internal/model"
	"gorm.io/gorm"
)

type Service struct {
	Repository Repository
	Validate   Validate
}

func NewService(dbconn *gorm.DB) Service {
	return Service{
		Repository: NewRepository(dbconn),
		Validate:   NewValidate(),
	}
}

func (service Service) Create(req model.RequestCreateItem) (model.Item, error) {
	now := time.Now()

	data := model.Item{
		Title:       req.Title,
		Amount:      req.Amount,
		Quantity:    req.Quantity,
		Status:      constant.ItemPendingStatus,
		CreatedTime: now,
		UpdatedTime: now,
	}

	if err := service.Repository.Create(&data); err != nil {
		return model.Item{}, err
	}

	return data, nil
}

func (service Service) Find(query model.RequestFindItem) ([]model.Item, error) {
	return service.Repository.Find(query)
}

func (service Service) Replace(id uint, req model.RequestCreateItem) (model.Item, error) {
	// Find item
	result, err := service.Repository.GetByID(id)
	if err != nil {
		return model.Item{}, err
	}

	// Check status
	if err := service.Validate.UpdateItem(result.Status); err != nil {
		return model.Item{}, err
	}

	data := model.Item{
		ID:          result.ID,
		Title:       req.Title,
		Amount:      req.Amount,
		Quantity:    req.Quantity,
		Status:      result.Status,
		CreatedTime: result.CreatedTime,
		UpdatedTime: time.Now(),
	}

	if err := service.Repository.Replace(data); err != nil {
		return model.Item{}, err
	}

	return data, nil
}

func (service Service) UpdateStatus(id uint, status constant.ItemStatus) (model.Item, error) {
	// Find item
	data, err := service.Repository.GetByID(id)
	if err != nil {
		return model.Item{}, err
	}

	// Check status
	if err := service.Validate.ItemStatusFlow(data.Status, status); err != nil {
		return model.Item{}, err
	}

	data.Status = status
	data.UpdatedTime = time.Now()

	// Replace
	if err := service.Repository.Replace(data); err != nil {
		return model.Item{}, err
	}

	return data, nil
}

func (service Service) Delete(id uint) error {
	// Find item
	data, err := service.Repository.GetByID(id)
	if err != nil {
		return err
	}

	// Check status
	if err := service.Validate.DeleteItem(data.Status); err != nil {
		return err
	}

	return service.Repository.Delete(id)
}
