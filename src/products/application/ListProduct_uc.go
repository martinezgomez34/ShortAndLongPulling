package application

import "github.com/martinezgomez34/abarrote/src/products/domain"

type ListProductUseCase struct {
	repo domain.IProduct
}

func ListProductUC(repo domain.IProduct) *ListProductUseCase {
	return &ListProductUseCase{repo: repo}
}

func (s *ListProductUseCase) GetAll() ([]*domain.Product, error) {
	return s.repo.GetAll()
}
