package service

import (
	"context"
	"errors"

	"github.com/fofow/backend-go/internal/model"
)

func (s *service) ProductAdjust(ctx context.Context, product model.Product) (err error) {
	if product.Quantity != -1 {
		return errors.New("quantity not set minus one")
	}

	tmpProduct, err := s.repo.GetProductByID(ctx, product.ID)
	if err != nil {
		return err
	}

	product.Quantity = tmpProduct.Quantity - 1
	err = s.repo.UpdateProduct(ctx, &product)
	if err != nil {
		return err
	}

	return err
}
