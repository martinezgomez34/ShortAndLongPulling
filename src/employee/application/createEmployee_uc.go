package application

import "github.com/martinezgomez34/abarrote/src/employee/domain"

type CreateEmployeeUseCase struct {
	repo domain.IEmployee
}

func CreateEmployee(repo domain.IEmployee) *CreateEmployeeUseCase {
	return &CreateEmployeeUseCase{repo: repo}
}

func (s *CreateEmployeeUseCase) Create(employee *domain.Employee) error {
	return s.repo.Save(employee)
}