package model

type RegisterAstraRequest struct {
	ID       int32  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
	Company  string `json:"company" db:"company"`
	UUID     string `json:"uuid" db:"uuid"`
	QRBase64 string `json:"qr_base64"`
}

type RegisterAstraResponse struct {
	Message string               `json:"message"`
	Data    RegisterAstraRequest `json:"data"`
}

type Astra struct {
	ID       int32  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
	Company  string `json:"company" db:"company"`
	UUID     string `json:"uuid" db:"uuid"`
	QRBase64 string `json:"qr_base64"`
}

type GetDatasAstraResponse struct {
	Message string  `json:"message"`
	Data    []Astra `json:"data"`
}

type GetDataAstraResponse struct {
	Message string `json:"message"`
	Data    Astra  `json:"data"`
}

type SearchAstraRequest struct {
	Search string `json:"search"`
}

type SearchAstraResponse struct {
	Message string               `json:"message"`
	Data    RegisterAstraRequest `json:"data"`
}

type GetDataIDsAstraResponse struct {
	Message string  `json:"message"`
	Data    []int32 `json:"data"`
}
