package application

import "github.com/martinezgomez34/abarrote/src/employee/domain"

type DeleteEmployeeUseCase struct {
	repo domain.IEmployee
}

func DeleteEmployee(repo domain.IEmployee) *DeleteEmployeeUseCase{
	return &DeleteEmployeeUseCase{repo: repo}
}

func (s *DeleteEmployeeUseCase) Delete(id int32) error{
	return s.repo.Delete(id)
}