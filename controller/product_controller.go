package controller

import "github.com/gin-gonic/gin"

type ProductController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
}
