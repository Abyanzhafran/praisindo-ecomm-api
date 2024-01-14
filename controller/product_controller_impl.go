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

	// Mapping required error data
	productResponseError := response.ProductErrorResponse{
		CorrelationID: uuid.NewString(),
		Success:       false,
		Error:         "",
		Tin:           time.Now(),
		Tout:          time.Now(),
		Data:          nil,
	}
	// Retrieve products from the repository
	products, err := c.ProductRepo.GetAll(ctx.Context())

	// Error handling
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	// Dereference each element in the slice
	var dereferencedProducts []domain.Product
	for _, p := range products {
		dereferencedProducts = append(dereferencedProducts, *p)
	}

	productListResponse := response.ProductListResponse{
		CorrelationID: uuid.NewString(),
		Success:       true,
		Error:         "",
		Tin:           time.Now(),
		Tout:          time.Now(),
		Data: response.ProductList{
			List:       dereferencedProducts,
			TotalItems: len(dereferencedProducts),
			TotalPages: 1,
			Page:       1,
			PageSize:   len(dereferencedProducts),
			Start:      time.Now(),
			Finish:     time.Now(),
			Duration:   "some-duration", // Add some duration
		},
	}

	return ctx.JSON(productListResponse)
}

func (c *ProductControllerImpl) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	products, err := c.ProductRepo.GetById(ctx.Context(), id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	singleProductResponse := response.ProductSingleResponse{
		CorrelationID: "some-correlation-id",
		Success:       true,
		Error:         "some error message",
		Tin:           time.Now(),
		Tout:          time.Now(),
		Data:          *products,
	}

	return ctx.Status(http.StatusOK).JSON(singleProductResponse)
}

func (c *ProductControllerImpl) Create(ctx *fiber.Ctx) error {
	var product domain.Product

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Check your input data",
			"error":   err.Error(),
		})
	}

	// generate uuid for the book
	generatedUuid, err := uuid.NewRandom()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error generating uuid",
		})
	}

	product.IDProduct = generatedUuid.String()

	if err := c.ProductRepo.Add(ctx.Context(), &product); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
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
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Id not found",
		})
	}

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Check your input data",
			"error":   err.Error(),
		})
	}

	if err := c.ProductRepo.Update(ctx.Context(), product); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
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

	var product domain.Product

	err := c.ProductRepo.Delete(ctx.Context(), id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Id not found",
		})
	}

	if err := c.ProductRepo.Delete(ctx.Context(), id); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
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
