package repository

import (
	"context"

	"gitlab.com/msstoci/popow-api/internal/model"
)

func (r *repository) GetDataSinarmas(ctx context.Context) (res []model.Sinarmas, err error) {
	err = r.db.GetSlave().SelectContext(
		ctx,
		&res,
		`
			SELECT 
				id, 
				name, 
				phone, 
				email, 
				attendance, 
				attendance_count, 
				product_name, 
				uuid 
			FROM 
				user
		`,
	)

	return res, err
}
