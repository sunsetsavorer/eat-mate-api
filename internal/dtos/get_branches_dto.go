package dtos

type GetBranchesDTO struct {
	Page  int
	Limit int
	Query *string
}

func (dto GetBranchesDTO) GetPage() int {
	return dto.Page
}

func (dto GetBranchesDTO) GetLimit() int {
	return dto.Limit
}

func (dto GetBranchesDTO) GetQuery() *string {
	return dto.Query
}
