package application

import (
	"github.com/martinezgomez34/abarrote/src/products/domain"
	"sort"
)

type NotifyLowStockUseCase struct {
	repo domain.IProduct
}

func NewNotifyLowStockUC(repo domain.IProduct) *NotifyLowStockUseCase {
	return &NotifyLowStockUseCase{repo: repo}
}

func (s *NotifyLowStockUseCase) GetLowStockProducts() ([]*domain.Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var lowStockProducts []*domain.Product
	for _, product := range products {
		if product.Amount < 15 {
			lowStockProducts = append(lowStockProducts, product)
		}
	}

	sort.Slice(lowStockProducts, func(i, j int) bool {
		return lowStockProducts[i].Amount < lowStockProducts[j].Amount
	})

	return lowStockProducts, nil
}
