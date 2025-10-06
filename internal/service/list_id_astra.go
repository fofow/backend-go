package service

import (
	"context"
)

func (s *service) ListIDsAstra(ctx context.Context) (ids []int32, err error) {

	data, err := s.repo.GetIDsAstra(ctx)
	if err != nil {
		return ids, err
	}

	return data, err
}
