package repositories

import (
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db/models"
)

type GroupRepository struct {
	db *db.Db
}

func NewGroupRepository(db *db.Db) *GroupRepository {

	return &GroupRepository{db}
}

func (r GroupRepository) Create(entity entities.GroupEntity) (entities.GroupEntity, error) {

	var group models.GroupModel

	group.FromEntity(entity)

	tx := r.db.Client.Begin()

	err := tx.Create(&group).Error
	if err != nil {
		tx.Rollback()
		return entities.GroupEntity{}, exceptions.NewRepositoryError(err)
	}

	entityMembers := entity.GetMembers()
	members := make([]models.GroupMemberModel, len(entityMembers))

	for i, e := range entityMembers {
		e.SetGroupID(group.ID)
		members[i].FromEntity(e)
	}

	err = tx.Create(&members).Error
	if err != nil {
		tx.Rollback()
		return entities.GroupEntity{}, exceptions.NewRepositoryError(err)
	}

	entityBranchOptions := entity.GetBranchOptions()
	branchOptions := make([]models.BranchModel, len(entityBranchOptions))

	for i, e := range entityBranchOptions {
		branchOptions[i].FromEntity(e)
	}

	err = tx.Model(&group).Association("BranchOptions").Append(&branchOptions)
	if err != nil {
		tx.Rollback()
		return entities.GroupEntity{}, exceptions.NewRepositoryError(err)
	}

	tx.Commit()

	return group.ToEntity(), nil
}
