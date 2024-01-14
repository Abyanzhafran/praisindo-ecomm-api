package controller

import (
	"dev/models/domain"
	"dev/models/http/response"
	"dev/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		Data: response.ProductList{
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

	singleProductResponse := response.ProductSingleResponse{
		CorrelationID: "some-correlation-id",
		Success:       true,
		Error:         "some error message",
		Tin:           time.Now(),
		Tout:          time.Now(),
		Data:          *products,
	}

	ctx.JSON(http.StatusOK, singleProductResponse)
}

func (c *ProductControllerImpl) Create(ctx *gin.Context) {
	var product domain.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Check your input data",
			"error":   err.Error(),
		})
		return
	}

	// generate uuid for the book
	generatedUuid, err := uuid.NewRandom()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Error generating uuid",
		})
	}

	product.IDProduct = generatedUuid.String()

	if err := c.ProductRepo.Add(ctx, &product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	singleProductResponse := response.ProductSingleResponse{
		CorrelationID: uuid.NewString(),
		Success:       true,
		Error:         "",
		Tin:           time.Now(),
		Tout:          time.Now(),
		Data:          product,
	}

	ctx.JSON(http.StatusOK, singleProductResponse)
}
