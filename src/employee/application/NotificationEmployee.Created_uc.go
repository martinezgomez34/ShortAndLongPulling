package application

import (
	"time"

	"github.com/martinezgomez34/abarrote/src/employee/domain"
)

type NotifyEmployeeCreatedUseCase struct {
	notificationChannel chan *domain.Employee
}

func NewNotifyEmployeeCreatedUC() *NotifyEmployeeCreatedUseCase {
	return &NotifyEmployeeCreatedUseCase{
		notificationChannel: make(chan *domain.Employee),
	}
}

func (s *NotifyEmployeeCreatedUseCase) WaitForNewEmployee() (*domain.Employee, error) {
	select {
	case employee := <-s.notificationChannel:
		return employee, nil
	case <-time.After(30 * time.Second):
		return nil, nil
	}
}

func (s *NotifyEmployeeCreatedUseCase) NotifyEmployeeCreated(employee *domain.Employee) {
	s.notificationChannel <- employee
}
