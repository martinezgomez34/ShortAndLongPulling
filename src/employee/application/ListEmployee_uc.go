package application

import "github.com/martinezgomez34/abarrote/src/employee/domain"

type ListEmployeeUseCase struct {
	repo domain.IEmployee
}

func ListEmployee(repo domain.IEmployee) *ListEmployeeUseCase {
	return &ListEmployeeUseCase{repo: repo}
}

func (s *ListEmployeeUseCase) GetAll() ([]*domain.Employee, error) {
	return s.repo.GetAll()
}