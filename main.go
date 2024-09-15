package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var products []Product

func main() {
	router := gin.Default()

	// Create a new product
	router.POST("/products", func(c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		products = append(products, newProduct)
		c.JSON(http.StatusCreated, newProduct)
	})

	// Get all products
	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	// Get a specific product by ID
	router.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, product := range products {
			if product.ID == id {
				c.JSON(http.StatusOK, product)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	// Update a product by ID
	router.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")

		var updatedProduct Product
		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, product := range products {
			if product.ID == id {
				products[i] = updatedProduct
				c.JSON(http.StatusOK, updatedProduct)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	// Delete a product by ID
	router.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")

		for i, product := range products {
			if product.ID == id {
				products = append(products[:i], products[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	router.Run(":8080")
}
