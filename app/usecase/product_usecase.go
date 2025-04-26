package usecase

import (
	"log"
	"simple-go-api/app/models"
	"simple-go-api/app/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]models.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) InsertProduct(product models.Product) (models.Product, error) {
	productId, err := pu.repository.InsertProduct(product)
	if err != nil {
		log.Fatalf("Unable to fetch product ID: %v", err)
		return models.Product{}, err
	}

	product.ID = productId

	return product, nil

}
