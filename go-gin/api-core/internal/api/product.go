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

type ProductImage struct {
	ImageName string `json:"image_name" binding:"required"`
	URL       string `json:"url" binding:"required,url,image"`
}

type CreateProductRequest struct {
	Name        string         `json:"name" binding:"required,min=3,max=100"`
	Description string         `json:"description" binding:"required,min=10"`
	Price       float64        `json:"price" binding:"required,gt=0"`
	Category    string         `json:"category" binding:"required,oneof=electronics books clothing home"`
	Display     *bool          `json:"display" binding:"omitempty"`
	Images      []ProductImage `json:"images" binding:"required,dive,required"`
	Tags        []string       `json:"tags" binding:"omitempty,dive,min=2,max=30"`
	Metadata    map[string]any `json:"metadata" binding:"omitempty"`
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

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	if req.Display == nil {
		defaultDisplay := true
		req.Display = &defaultDisplay
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Create product - v1",
		"data":    req,
	})
}
