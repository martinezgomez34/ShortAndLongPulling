package application

import (
	"sync"

)

type EmployeeAction struct {
	ID     int32
	Action string
}

type NotifyEmployeeActionUseCase struct {
	mu              sync.Mutex
	actionRegistry  []EmployeeAction
}

func NewNotifyEmployeeActionUC() *NotifyEmployeeActionUseCase {
	return &NotifyEmployeeActionUseCase{}
}

func (s *NotifyEmployeeActionUseCase) RegisterAction(id int32, action string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.actionRegistry = append(s.actionRegistry, EmployeeAction{
		ID:     id,
		Action: action,
	})
}

func (s *NotifyEmployeeActionUseCase) GetActions() []EmployeeAction {
	s.mu.Lock()
	defer s.mu.Unlock()
	actions := s.actionRegistry
	s.actionRegistry = nil
	return actions
}
