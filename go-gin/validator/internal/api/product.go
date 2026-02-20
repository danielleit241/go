package api

import (
	"github.com/danielleit241/utils"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

type GetByProductIDParam struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (ctrl *ProductController) GetProducts(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all products - v1",
	})
}

func (ctrl *ProductController) SearchProducts(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Search products - v1",
	})
}

func (ctrl *ProductController) GetProductBySlug(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get product by slug - v1",
	})
}

func (ctrl *ProductController) GetProductsByCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get products by category - v1",
	})
}

func (ctrl *ProductController) GetProductByID(c *gin.Context) {
	var param GetByProductIDParam
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(400, utils.HandleValidationError(err))
		return
	}
	c.JSON(200, gin.H{
		"message": "Get product by ID - v1",
	})
}
