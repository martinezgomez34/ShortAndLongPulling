package application

import "github.com/martinezgomez34/abarrote/src/employee/domain"

type GetByIDEmployeeUseCase struct {
	repo domain.IEmployee
}

func GetByIDEmployee(repo domain.IEmployee) *GetByIDEmployeeUseCase {
	return &GetByIDEmployeeUseCase{repo: repo}
}

func (s *GetByIDEmployeeUseCase) GetByID(id int32) (*domain.Employee, error) {
	return s.repo.GetByID(id)
}
