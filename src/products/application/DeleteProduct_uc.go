package application

import "github.com/martinezgomez34/abarrote/src/products/domain"

type DeleteProductUseCase struct {
	repo domain.IProduct
}

func DeleteProductUC(repo domain.IProduct) *DeleteProductUseCase {
	return &DeleteProductUseCase{repo: repo}
}

func (s *DeleteProductUseCase) Delete(id int32) error {
	return s.repo.Delete(id)
}