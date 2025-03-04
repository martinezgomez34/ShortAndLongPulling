package application

import "github.com/martinezgomez34/abarrote/src/employee/domain"

type EditEmployeeUseCase struct {
	repo domain.IEmployee
}

func EditEmployee(repo domain.IEmployee) *EditEmployeeUseCase{
	return &EditEmployeeUseCase{repo: repo}
}

func (s *EditEmployeeUseCase) Update(employee *domain.Employee) error{
	return s.repo.Update(employee)
}