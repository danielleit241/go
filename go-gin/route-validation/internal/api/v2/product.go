package v2handler

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Slug     string  `json:"slug"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

var products = []Product{
	{ID: 1, Name: "Product A", Slug: "product-a", Price: 19.99, Category: "Electronics"},
	{ID: 2, Name: "Product B", Slug: "product-b", Price: 29.99, Category: "Clothing"},
	{ID: 3, Name: "Product C", Slug: "product-c", Price: 39.99, Category: "Books"},
}

var slugRegex = `^[a-z0-9]+(?:-[a-z0-9]+)*$`
var searchRegex = `^[a-zA-Z0-9\s]{6,99}$`

var validCategories = map[string]bool{
	"Electronics": true,
	"Clothing":    true,
	"Books":       true,
}

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (h *ProductHandler) GetProductsV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (h *ProductHandler) SearchProductsV2(c *gin.Context) {
	search := c.Query("query")
	if search == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter 'query' is required",
		})
		return
	}
	matched, err := regexp.MatchString(searchRegex, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter 'query' must be 6-99 characters long and contain only alphanumeric characters and spaces",
		})
		return
	}
	limitParam := c.DefaultQuery("limit", "10") // if limit is not provided, default to 10
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter 'limit' must be a positive integer",
		})
		return
	}
	var results []Product
	re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(search))
	for _, product := range products {
		// Case-insensitive search for the product name
		if re.MatchString(product.Name) || re.MatchString(product.Slug) {
			results = append(results, product)
			if len(results) >= limit {
				break
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"products": results,
	})
}

func (h *ProductHandler) GetProductBySlugV2(c *gin.Context) {
	slug := c.Param("slug")

	matched, err := regexp.MatchString(slugRegex, slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid slug format",
		})
		return
	}

	for _, product := range products {
		if product.Slug == slug {
			c.JSON(http.StatusOK, gin.H{
				"product": product,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Product not found",
	})
}

func (h *ProductHandler) GetProductsByCategoryV2(c *gin.Context) {
	category := c.Param("category")
	if !validCategories[category] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid category",
		})
		return
	}
	var filteredProducts []Product
	for _, product := range products {
		if product.Category == category {
			filteredProducts = append(filteredProducts, product)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"products": filteredProducts,
	})
}

func (h *ProductHandler) CreateProductV2(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"product": "new product created (v2)",
	})
}

func (h *ProductHandler) UpdateProductV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"product": "product updated (v2)",
	})
}

func (h *ProductHandler) DeleteProductV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "product deleted (v2)",
	})
}
