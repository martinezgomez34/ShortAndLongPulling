package application

import "github.com/martinezgomez34/abarrote/src/products/domain"

type GetByIDProductUseCase struct {
	repo domain.IProduct
}

func GetByIDProductUC(repo domain.IProduct) *GetByIDProductUseCase {
	return &GetByIDProductUseCase{repo: repo}
}

func (s *GetByIDProductUseCase) GetByID(id int32) (*domain.Product, error) {
	return s.repo.GetByID(id)
}
