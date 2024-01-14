package controller

import (
	"dev/models/domain"
	"dev/models/http/response"
	"dev/repository"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductControllerImpl struct {
	ProductRepo repository.ProductRepository
}

func NewProductController(ProductRepo repository.ProductRepository) ProductController {
	return &ProductControllerImpl{ProductRepo: ProductRepo}
}

func (c *ProductControllerImpl) FindAll(ctx *fiber.Ctx) error {
	start := time.Now()

	// Retrieve products from the repository
	products, err := c.ProductRepo.GetAll(ctx.Context())

	// Error handling
	if err != nil {
		// Mapping required error data
		productResponseError := response.ProductErrorResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          nil,
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	// Dereference each element in the slice
	var dereferencedProducts []domain.Product
	for _, p := range products {
		dereferencedProducts = append(dereferencedProducts, *p)
	}

	finish := time.Now()

	// Benchmark the query duration
	// idk if it's true or not for the bencmark
	duration := finish.Sub(start).String()

	productListResponse := response.ProductListResponse{
		CorrelationID: uuid.NewString(),
		Success:       true,
		Error:         "",
		Tin:           start,
		Tout:          finish,
		Data: response.ProductList{
			List:       dereferencedProducts,
			TotalItems: len(dereferencedProducts),
			TotalPages: 1,
			Page:       1,
			PageSize:   len(dereferencedProducts),
			Start:      start,
			Finish:     finish,
			Duration:   duration,
		},
	}

	return ctx.JSON(productListResponse)
}

func (c *ProductControllerImpl) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	products, err := c.ProductRepo.GetById(ctx.Context(), id)

	if err != nil {
		productResponseError := response.ProductErrorResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          nil,
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	singleProductResponse := response.ProductSingleResponse{
		CorrelationID: uuid.NewString(),
		Success:       true,
		Error:         "",
		Tin:           time.Now(),
		Tout:          time.Now(),
		Data:          *products,
	}

	return ctx.Status(http.StatusOK).JSON(singleProductResponse)
}

func (c *ProductControllerImpl) Create(ctx *fiber.Ctx) error {
	var product domain.Product

	// Request data should in accordance with product model(domain)
	if err := ctx.BodyParser(&product); err != nil {
		productResponseError := response.ProductErrorResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          nil,
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	// Generate uuid for the product
	generatedUuid, err := uuid.NewRandom()
	if err != nil {
		productResponseError := response.ProductErrorResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          nil,
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	product.IDProduct = generatedUuid.String()

	if err := c.ProductRepo.Add(ctx.Context(), &product); err != nil {
		productResponseError := response.ProductErrorResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          nil,
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	singleProductResponse := response.ProductSingleResponse{
		CorrelationID: uuid.NewString(),
		Success:       true,
		Error:         "",
		Tin:           time.Now(),
		Tout:          time.Now(),
		Data:          product,
	}

	return ctx.Status(http.StatusOK).JSON(singleProductResponse)
}

func (c *ProductControllerImpl) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	product, err := c.ProductRepo.GetById(ctx.Context(), id)
	if err != nil {
		productResponseError := response.ProductErrorResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          nil,
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	if err := ctx.BodyParser(&product); err != nil {
		productResponseError := response.ProductErrorResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          nil,
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	if err := c.ProductRepo.Update(ctx.Context(), product); err != nil {
		productResponseError := response.ProductErrorResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          nil,
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	singleProductResponse := response.ProductSingleResponse{
		CorrelationID: uuid.NewString(),
		Success:       true,
		Error:         "",
		Tin:           time.Now(),
		Tout:          time.Now(),
		Data:          *product,
	}

	return ctx.Status(http.StatusOK).JSON(singleProductResponse)
}

func (c *ProductControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	_, err := c.ProductRepo.GetById(ctx.Context(), id)
	if err != nil {
		productResponseError := response.ProductErrorResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          nil,
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	if err := c.ProductRepo.Delete(ctx.Context(), id); err != nil {
		productResponseError := response.ProductErrorResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          nil,
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	singleProductResponse := response.ProductSingleResponse{
		CorrelationID: uuid.NewString(),
		Success:       true,
		Error:         "",
		Tin:           time.Now(),
		Tout:          time.Now(),
		Data:          nil,
	}

	return ctx.Status(http.StatusOK).JSON(singleProductResponse)
}
