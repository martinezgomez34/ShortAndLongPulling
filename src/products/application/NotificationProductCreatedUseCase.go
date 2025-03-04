package application

import (
	"time"
	"github.com/martinezgomez34/abarrote/src/products/domain"
)

type NotifyProductCreatedUseCase struct {
	repo domain.IProduct
}

func NewNotifyProductCreatedUC(repo domain.IProduct) *NotifyProductCreatedUseCase {
	return &NotifyProductCreatedUseCase{repo: repo}
}

func (s *NotifyProductCreatedUseCase) WaitForNewProduct() ([]*domain.Product, error) {
	select {
	case <-WaitForProductCreated():
		products, err := s.repo.GetAll()
		if err != nil {
			return nil, err
		}

		return products, nil

	case <-time.After(30 * time.Second):
		return nil, nil
	}
}
