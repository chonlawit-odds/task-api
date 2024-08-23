package item

import (
	"github.com/chonlawit-odds/task-api/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	Database *gorm.DB
}

func NewRepository(dbconn *gorm.DB) Repository {
	return Repository{
		Database: dbconn,
	}
}

func (repo Repository) Create(data *model.Item) error {
	return repo.Database.Create(&data).Error
}

func (repo Repository) GetByID(id uint) (model.Item, error) {
	var result model.Item

	if err := repo.Database.First(&result, id).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (repo Repository) Find(query model.RequestFindItem) ([]model.Item, error) {
	var results []model.Item

	tx := repo.Database

	if statuses := query.Statuses; len(statuses) > 0 {
		tx = tx.Where("status IN ?", query.Statuses)
	}

	if err := tx.Find(&results).Error; err != nil {
		return results, err
	}

	return results, nil
}

func (repo Repository) Replace(data model.Item) error {
	return repo.Database.Model(&data).Updates(data).Error
}

func (repo Repository) Delete(id uint) error {
	return repo.Database.Delete(&model.Item{}, id).Error
}
