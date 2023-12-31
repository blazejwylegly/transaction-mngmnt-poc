package service

import (
	"github.com/blazejwylegly/product-service/product/data"
	"github.com/blazejwylegly/product-service/product/models"
)

type ProductService struct {
	productRepo data.ProductRepository
}

func (service ProductService) FindAll() []*models.ProductDto {
	var products []*models.ProductDto
	for _, entity := range service.productRepo.FindAll() {
		productDto := &models.ProductDto{
			ProductId:   entity.ProductID,
			Name:        entity.Name,
			Price:       entity.Price,
			Description: entity.Description,
			Quantity:    entity.Quantity,
		}
		products = append(products, productDto)
	}
	return products
}

func (service ProductService) SaveAll(products []*models.ProductDto) {
	for _, product := range products {
		productEntity := &models.Product{
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Quantity:    product.Quantity,
		}
		service.productRepo.SaveOrUpdate(productEntity)
	}
}

func (service ProductService) SaveOrUpdate(product *models.ProductDto) {
	productEntity := &models.Product{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    product.Quantity,
	}
	service.productRepo.SaveOrUpdate(productEntity)
}

func (service ProductService) FindById(productId int) *models.ProductDto {
	entity := service.productRepo.FindById(productId)
	return &models.ProductDto{
		ProductId:   entity.ProductID,
		Name:        entity.Name,
		Price:       entity.Price,
		Description: entity.Description,
		Quantity:    entity.Quantity,
	}
}

func (service ProductService) UpdateQuantity(productId int, body struct {
	Quantity int `json:"quantity"`
}) {
	service.productRepo.UpdateQuantity(productId, body.Quantity)
}

func NewProductService(productRepo *data.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: *productRepo,
	}
}
