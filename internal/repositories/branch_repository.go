package repositories

import (
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db/models"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases/branch"
)

type BranchRepository struct {
	db *db.Db
}

func NewBranchRepository(db *db.Db) *BranchRepository {

	return &BranchRepository{db}
}

func (r BranchRepository) GetList(filter branch.BranchesFilter) ([]entities.BranchEntity, int64, error) {

	var branches []models.BranchModel
	var count int64

	query := r.db.Client.
		Model(&models.BranchModel{}).
		Joins("Brand")

	if filter.Query != nil {
		query = query.Where("\"Brand\".\"name\" ILIKE ?", "%"+*filter.Query+"%")
	}

	err := query.
		Count(&count).
		Error
	if err != nil {
		return nil, 0, exceptions.NewRepositoryError(err)
	}

	offset := (filter.Page - 1) * filter.Limit

	err = query.
		Limit(filter.Limit).
		Offset(offset).
		Preload("Brand").
		Find(&branches).
		Error
	if err != nil {
		return nil, 0, exceptions.NewRepositoryError(err)
	}

	var result []entities.BranchEntity = make([]entities.BranchEntity, len(branches))

	for i, e := range branches {
		result[i] = e.ToEntity()
	}

	return result, count, nil
}
