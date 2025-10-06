package service

import (
	"context"

	"github.com/fofow/backend-go/internal/model"
)

func (s *service) GetProductsActive(ctx context.Context) (products []model.Product, err error) {

	products, err = s.repo.GetProducts(ctx)
	if err != nil {
		return products, err
	}

	var tmpProduct []model.Product
	for _, product := range products {
		product.IsSelected = true
		if product.Quantity == 0 {
			product.IsSelected = false
		}

		tmpProduct = append(tmpProduct, product)
	}

	return tmpProduct, err
}
