package service

import (
	"context"
)

func (s *service) DeleteProductByID(ctx context.Context, id int32) (err error) {

	err = s.repo.SoftDeleteProductByID(ctx, id)
	if err != nil {
		return err
	}

	return err
}
