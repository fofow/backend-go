package model

type Product struct {
	ID         int32  `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Quantity   int32  `json:"quantity" db:"quantity"`
	IsSelected bool   `json:"is_selected,omitempty"`
}

type GetProductsResponse struct {
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}
