package application

import (
	"time"
	"github.com/martinezgomez34/abarrote/src/products/domain"
)

type NotifyProductActionUseCase struct {
	repo domain.IProduct
}

func NewNotifyProductActionUC(repo domain.IProduct) *NotifyProductActionUseCase {
	return &NotifyProductActionUseCase{repo: repo}
}

func (s *NotifyProductActionUseCase) WaitForAction() ([]*domain.Product, error) {
	select {
	case <-WaitForProductAction():
		products, err := s.repo.GetAll()
		if err != nil {
			return nil, err
		}

		return products, nil

	case <-time.After(30 * time.Second):
		return nil, nil
	}
}
