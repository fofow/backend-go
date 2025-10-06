package model

type Data struct {
	ID           int32  `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Email        string `json:"email" db:"email"`
	Telephone    string `json:"phone" db:"phone"`
	Company      string `json:"company" db:"company"`
	QRBase64     string `json:"qr_base64"`
	RegisterDate string `json:"register_date" db:"register_date"`
}

type Response struct {
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Login struct {
	Email string `json:"email"`
}

type RegisterSinarmasRequest struct {
	ID              int32  `json:"id" db:"id"`
	Name            string `json:"name" db:"name"`
	Email           string `json:"email" db:"email"`
	Telephone       string `json:"phone" db:"phone"`
	Attendance      string `json:"attendance" db:"attendance"`
	AttendanceCount int32  `json:"attendance_count" db:"attendance_count"`
	ProductName     string `json:"product_name" db:"product_name"`
	RegisterDate    string `json:"register_date" db:"register_date"`
	UUID            string `json:"uuid" db:"uuid"`
	QRBase64        string `json:"qr_base64"`
}

type RegisterSinarmasResponse struct {
	Message string                  `json:"message"`
	Data    RegisterSinarmasRequest `json:"data"`
}

type SearchSinarmasRequest struct {
	Phone string `json:"phone" db:"phone"`
}

type SearchSinarmasResponse struct {
	Message string                  `json:"message"`
	Data    RegisterSinarmasRequest `json:"data"`
}

type Sinarmas struct {
	ID              int32  `json:"id" db:"id"`
	Name            string `json:"name" db:"name"`
	Email           string `json:"email" db:"email"`
	Telephone       string `json:"phone" db:"phone"`
	Attendance      string `json:"attendance" db:"attendance"`
	AttendanceCount int32  `json:"attendance_count" db:"attendance_count"`
	ProductName     string `json:"product_name" db:"product_name"`
	RegisterDate    string `json:"register_date" db:"register_date"`
	UUID            string `json:"uuid" db:"uuid"`
	QRBase64        string `json:"qr_base64"`
}

type GetDatasSinarmasResponse struct {
	Message string     `json:"message"`
	Data    []Sinarmas `json:"data"`
}

type GetDataSinarmasResponse struct {
	Message string   `json:"message"`
	Data    Sinarmas `json:"data"`
}
