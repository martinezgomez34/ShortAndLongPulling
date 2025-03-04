package application

import "github.com/martinezgomez34/abarrote/src/products/domain"

type CreateProductUseCase struct {
	repo domain.IProduct
}

func NewProductUC(repo domain.IProduct) *CreateProductUseCase {
	return &CreateProductUseCase{repo: repo}
}

func (s *CreateProductUseCase) Create(product *domain.Product) error {
	return s.repo.Save(product)
}