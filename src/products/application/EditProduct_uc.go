package application

import "github.com/martinezgomez34/abarrote/src/products/domain"

type EditProductUseCase struct {
	repo domain.IProduct
}

func EditProductUC(repo domain.IProduct) *EditProductUseCase {
	return &EditProductUseCase{repo: repo}
}

func (s *EditProductUseCase) Update(product *domain.Product) error {
	return s.repo.Update(product)
}