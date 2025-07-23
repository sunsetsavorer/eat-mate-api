package repositories

import (
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db/models"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases/place"
)

type PlaceBranchRepository struct {
	db *db.Db
}

func NewPlaceBranchRepository(db *db.Db) *PlaceBranchRepository {

	return &PlaceBranchRepository{db}
}

func (r PlaceBranchRepository) GetList(filter place.PlaceBranchFilter) ([]entities.PlaceBranchEntity, int64, error) {

	var placeBranches []models.PlaceBranchModel
	var count int64

	query := r.db.Client.
		Model(&models.PlaceBranchModel{}).
		Joins("Place")

	if filter.Query != nil {
		query = query.Where("\"Place\".\"name\" ILIKE ?", "%"+*filter.Query+"%")
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
		Preload("Place").
		Find(&placeBranches).
		Error
	if err != nil {
		return nil, 0, exceptions.NewRepositoryError(err)
	}

	var result []entities.PlaceBranchEntity = make([]entities.PlaceBranchEntity, len(placeBranches))

	for i, e := range placeBranches {
		result[i] = e.ToEntity()
	}

	return result, count, nil
}
