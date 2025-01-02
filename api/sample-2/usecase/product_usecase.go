package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	// Repository
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

// Função que irá tratar a requisição
func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}
