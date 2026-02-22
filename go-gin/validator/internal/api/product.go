package api

import (
	"net/http"

	"github.com/danielleit241/utils"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

type GetByProductIDParam struct {
	ID string `uri:"id" binding:"uuid"`
}

type GetProductBySlugParam struct {
	Slug string `uri:"slug" binding:"slug,min=3,max=100"`
}

type GetProductsByCategoryParam struct {
	// Category string `uri:"category" binding:"required,category"` // Custom validator for category
	Category string `uri:"category" binding:"oneof=electronics books clothing home"` // Using built-in oneof validator
}

type SearchProductsQuery struct {
	Query string `form:"query" binding:"required,search,min=3,max=100"`
	Date  string `form:"date" binding:"omitempty,datetime=2006-01-02"`
	Limit int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (ctrl *ProductController) GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get all products - v1",
	})
}

func (ctrl *ProductController) SearchProducts(c *gin.Context) {
	var params SearchProductsQuery
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	if params.Limit == 0 {
		params.Limit = 10 // Default limit
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Search products - v1",
		"query":   params.Query,
		"limit":   params.Limit,
		"date":    params.Date,
	})
}

func (ctrl *ProductController) GetProductBySlug(c *gin.Context) {
	var param GetProductBySlugParam
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get product by slug - v1",
		"slug":    param.Slug,
	})
}

func (ctrl *ProductController) GetProductsByCategory(c *gin.Context) {
	var param GetProductsByCategoryParam
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Get products by category - v1",
		"category": param.Category,
	})
}

func (ctrl *ProductController) GetProductByID(c *gin.Context) {
	var param GetByProductIDParam
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get product by ID - v1",
		"id":      param.ID,
	})
}
