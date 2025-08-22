package repositories

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sunsetsavorer/eat-mate-api/internal/entities"
	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db"
	"github.com/sunsetsavorer/eat-mate-api/internal/infrastructure/db/models"
	"github.com/sunsetsavorer/eat-mate-api/internal/usecases/group"
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

func (r GroupRepository) GetList(filter group.GroupsFilter) ([]entities.GroupEntity, int64, error) {

	var groups []models.GroupModel
	var count int64

	query := r.db.Client.
		Model(models.GroupModel{}).
		Where("is_public = ?", true).
		Where("is_active = ?", true).
		Order("created_at desc")

	err := query.Count(&count).Error
	if err != nil {
		return nil, 0, exceptions.NewRepositoryError(err)
	}

	offset := (filter.Page - 1) * filter.Limit

	err = query.
		Offset(offset).
		Limit(filter.Limit).
		Preload("Branch.Brand").
		Preload("Members.User").
		Find(&groups).
		Error

	if err != nil {
		return nil, 0, exceptions.NewRepositoryError(err)
	}

	groupEntities := make([]entities.GroupEntity, len(groups))

	for i, e := range groups {
		groupEntities[i] = e.ToEntity()
	}

	return groupEntities, count, nil
}

func (r GroupRepository) GetByID(ID uuid.UUID) (entities.GroupEntity, error) {

	var group models.GroupModel

	err := r.db.Client.
		Model(&models.GroupModel{}).
		Where("is_active = ?", true).
		Preload("Branch.Brand").
		Preload("Members.User").
		Preload("BranchOptions.Brand").
		Preload("Votes").
		First(&group, ID).
		Error

	if err != nil {
		return entities.GroupEntity{}, exceptions.NewNotFoundError(fmt.Errorf("group with specified id wasn't found"))
	}

	return group.ToEntity(), nil
}

func (r GroupRepository) AddMember(entity entities.GroupMemberEntity) error {

	var member models.GroupMemberModel

	member.FromEntity(entity)

	err := r.db.Client.Create(&member).Error
	if err != nil {
		return exceptions.NewRepositoryError(err)
	}

	return nil
}

func (r GroupRepository) RemoveMember(entity entities.GroupMemberEntity) error {

	err := r.db.Client.
		Where("user_id = ?", entity.GetUserID()).
		Where("group_id = ?", entity.GetGroupID()).
		Delete(&models.GroupMemberModel{}).Error

	if err != nil {
		return exceptions.NewRepositoryError(err)
	}

	return nil
}

func (r GroupRepository) GetMemberByID(userID int64, groupID uuid.UUID) (entities.GroupMemberEntity, error) {

	var member models.GroupMemberModel

	err := r.db.Client.
		Model(&models.GroupMemberModel{}).
		Where("user_id = ?", userID).
		Where("group_id = ?", groupID).
		Preload("User").
		First(&member).
		Error

	if err != nil {
		return entities.GroupMemberEntity{}, exceptions.NewRepositoryError(err)
	}

	result := member.ToEntity()

	return result, nil
}

func (r GroupRepository) DeactivateByID(groupID uuid.UUID) error {

	err := r.db.Client.
		Model(&models.GroupModel{}).
		Where("id = ?", groupID).
		Update("is_active", false).
		Error

	if err != nil {
		return exceptions.NewRepositoryError(err)
	}

	return nil
}

func (r GroupRepository) SaveVote(entity entities.VoteEntity) error {

	var vote models.VoteModel

	rowsAffected := r.db.Client.
		Model(&models.VoteModel{}).
		Where("user_id", entity.GetUserID()).
		Where("group_id", entity.GetGroupID()).
		Update("branch_id", entity.GetBranchID()).
		RowsAffected

	if rowsAffected != 0 {
		return nil
	}

	vote.FromEntity(entity)

	err := r.db.Client.
		Create(&vote).
		Error

	if err != nil {
		return exceptions.NewRepositoryError(err)
	}

	return nil
}
