package dtos

type GetGroupsDTO struct {
	Page  int
	Limit int
}

func (dto GetGroupsDTO) GetPage() int {
	return dto.Page
}

func (dto GetGroupsDTO) GetLimit() int {
	return dto.Limit
}
