package controller

import (
	"dev/models/domain"
	"dev/models/http/response"
	"dev/repository"
	"math"
	"net/http"
	"strconv"
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
	// Get query param
	page := ctx.Query("page", "1")
	pageSize := ctx.Query("pageSize", "5")
	productName := ctx.Query("product-name")

	// If product query searched
	if productName != "" {
		products, err := c.ProductRepo.GetByProductName(ctx.Context(), productName)

		if err != nil {
			productResponseError := response.ProductSingleResponse{
				CorrelationID: uuid.NewString(),
				Success:       false,
				Error:         err.Error(),
				Tin:           time.Now(),
				Tout:          time.Now(),
				Data:          map[string]interface{}{},
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

	start := time.Now()

	// Atoi, Convert string to int
	parsedPage, err := strconv.Atoi(page)
	if err != nil || parsedPage < 1 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page parameter"})
	}

	parsedPageSize, err := strconv.Atoi(pageSize)
	if err != nil || parsedPageSize < 1 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid pageSize parameter"})
	}

	// Calculate offset based on page and pageSize
	// like max item in each page
	offset := (parsedPage - 1) * parsedPageSize

	products, totalCount, err := c.ProductRepo.GetPaginated(ctx.Context(), int64(parsedPageSize), offset)

	if err != nil {
		// Mapping required error data
		productResponseError := response.ProductSingleResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          map[string]interface{}{},
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	// Dereference each element in the slice
	var dereferencedProducts []domain.Product
	for _, p := range products {
		dereferencedProducts = append(dereferencedProducts, *p)
	}

	finish := time.Now()

	// Calculate total pages for paging
	// math.ceil func is always round up, similiar with JS
	totalPages := int(math.Ceil(float64(totalCount) / float64(parsedPageSize)))

	// Benchmark the query duration from start time
	duration := finish.Sub(start).String()

	productListResponse := response.ProductListResponse{
		CorrelationID: uuid.NewString(),
		Success:       true,
		Error:         "",
		Tin:           start,
		Tout:          finish,
		Data: response.ProductList{
			List:       dereferencedProducts,
			TotalItems: totalCount,
			TotalPages: totalPages,
			Page:       parsedPage,
			PageSize:   parsedPageSize,
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
		productResponseError := response.ProductSingleResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          map[string]interface{}{},
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
		productResponseError := response.ProductSingleResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          map[string]interface{}{},
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	// Generate uuid for the product
	generatedUuid, err := uuid.NewRandom()
	if err != nil {
		productResponseError := response.ProductSingleResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          map[string]interface{}{},
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	product.IDProduct = generatedUuid.String()

	if err := c.ProductRepo.Add(ctx.Context(), &product); err != nil {
		productResponseError := response.ProductSingleResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          map[string]interface{}{},
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
		productResponseError := response.ProductSingleResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          map[string]interface{}{},
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	if err := ctx.BodyParser(&product); err != nil {
		productResponseError := response.ProductSingleResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          map[string]interface{}{},
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	if err := c.ProductRepo.Update(ctx.Context(), product); err != nil {
		productResponseError := response.ProductSingleResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          map[string]interface{}{},
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
		productResponseError := response.ProductSingleResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          map[string]interface{}{},
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	if err := c.ProductRepo.Delete(ctx.Context(), id); err != nil {
		productResponseError := response.ProductSingleResponse{
			CorrelationID: uuid.NewString(),
			Success:       false,
			Error:         err.Error(),
			Tin:           time.Now(),
			Tout:          time.Now(),
			Data:          map[string]interface{}{},
		}

		return ctx.Status(http.StatusInternalServerError).JSON(productResponseError)
	}

	singleProductResponse := response.ProductSingleResponse{
		CorrelationID: uuid.NewString(),
		Success:       true,
		Error:         "",
		Tin:           time.Now(),
		Tout:          time.Now(),
		Data:          map[string]interface{}{},
	}

	return ctx.Status(http.StatusOK).JSON(singleProductResponse)
}
