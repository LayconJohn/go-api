package usecase

import "github/LayconJohn/go-api/internal/entity"

type CreateProductInputDto struct {
	Name  string
	Price float64
}

type CreateProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func (u *CreateProductUseCase) Execute(input CreateProductInputDto)
