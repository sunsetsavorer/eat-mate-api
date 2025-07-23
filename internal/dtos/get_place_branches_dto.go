package dtos

type GetPlaceBranchesDTO struct {
	Page  int
	Limit int
	Query *string
}

func (dto GetPlaceBranchesDTO) GetPage() int {
	return dto.Page
}

func (dto GetPlaceBranchesDTO) GetLimit() int {
	return dto.Limit
}

func (dto GetPlaceBranchesDTO) GetQuery() *string {
	return dto.Query
}
