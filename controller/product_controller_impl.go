package controller

import (
	"dev/models/domain"
	"dev/models/http/response"
	"dev/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductRepo repository.ProductRepository
}

func NewProductController(ProductRepo repository.ProductRepository) ProductController {
	return &ProductControllerImpl{ProductRepo: ProductRepo}
}

func (c *ProductControllerImpl) FindAll(ctx *gin.Context) {
	// Retrieve products from the repository
	products, err := c.ProductRepo.GetAll(ctx)
	if err != nil {
		// Handle the error and return an Internal Server Error response
		ctx.JSON(http.StatusInternalServerError, response.ProductListResponse{
			Success: false,
			Error:   "Internal Server Error",
		})
		return
	}

	// Dereference each element in the slice
	var dereferencedProducts []domain.Product
	for _, p := range products {
		dereferencedProducts = append(dereferencedProducts, *p)
	}

	productListResponse := response.ProductListResponse{
		CorrelationID: "some-correlation-id", // You may want to generate a correlation ID dynamically
		Success:       true,
		Error:         "",
		Tin:           time.Now(),
		Tout:          time.Now(),
		Data: response.ProductDetail{
			List:       dereferencedProducts,
			TotalItems: len(dereferencedProducts),
			TotalPages: 1,                         // You may need to calculate the total pages based on your pagination logic
			Page:       1,                         // You may need to get the current page from the request
			PageSize:   len(dereferencedProducts), // Assuming no pagination, adjust accordingly
			Start:      time.Now(),
			Finish:     time.Now(),
			Duration:   "some-duration", // You may want to calculate the duration
		},
	}

	// Return the ProductListResponse as JSON
	ctx.JSON(http.StatusOK, productListResponse)
}

func (c *ProductControllerImpl) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	products, err := c.ProductRepo.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
