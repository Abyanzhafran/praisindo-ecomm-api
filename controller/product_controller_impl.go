package controller

import (
	"dev/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductRepo repository.ProductRepository
}

func NewProductController(ProductRepo repository.ProductRepository) ProductController {
	return &ProductControllerImpl{ProductRepo: ProductRepo}
}

func (c *ProductControllerImpl) FindAll(ctx *gin.Context) {
	products, err := c.ProductRepo.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, products)
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
