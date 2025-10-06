package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/fofow/backend-go/internal/handler"
	"github.com/fofow/backend-go/internal/repository"
	"github.com/fofow/backend-go/internal/service"
	"github.com/fofow/backend-go/pkg/database"
)

func main() {

	db := database.New(database.DBConfig{
		SlaveDSN:      "popow:pPdtMXjHhgEI@tcp(localhost:3306)/registrasi?parseTime=true",
		MasterDSN:     "popow:pPdtMXjHhgEI@tcp(localhost:3306)/registrasi?parseTime=true",
		RetryInterval: 10,
		MaxIdleConn:   10,
		MaxConn:       10,
	}, database.DriverMySQL)

	repo := repository.New(db)

	svc := service.New(repo)

	delivery := handler.New(svc)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
    AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
}))

	// Routes
	e.POST("/login", delivery.GetDataByEmail)
	e.POST("/register", delivery.Register)
	e.POST("/sinarmas/register", delivery.RegisterSinarmas)
	e.POST("/astra/register", delivery.RegisterAstra)
	e.GET("/sinarmas/search", delivery.SearchSinarmas)
	e.GET("/sinarmas", delivery.ListDataSinarmas)
	e.POST("/sinarmas/:id/activated", delivery.UpdateAttendance)

	e.GET("/astra", delivery.ListDataAstra)
	e.GET("/astra/search", delivery.SearchAstra)
	e.POST("/astra/:id/activated", delivery.UpdateAttendanceAstra)
	e.GET("/astra/activated-list", delivery.ListIDAstra)
	e.POST("/astra/:id/winner", delivery.UpdateWinnerAstra)

	e.GET("/product", delivery.GetProduct)
	e.POST("/product", delivery.CreateProduct)
	e.PATCH("/product/:id", delivery.EditProduct)
	e.GET("/product/:id", delivery.GetProductByID)
	e.DELETE("/product/:id", delivery.DeleteProductByID)

	e.GET("/wheel/products", delivery.GetProductsActive)
	e.POST("/wheel/product-adjust", delivery.ProductAdjust)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {

	type Data struct {
		Name      string
		Email     string
		Telephone string
		Company   string
	}

	type Response struct {
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}

	var res Response

	var data Data

	data.Name = "test"
	data.Email = "test"
	data.Telephone = "test"
	data.Company = "test"

	res.Data = data

	return c.JSON(http.StatusOK, res)
}
