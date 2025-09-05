package handlers

import (
	"E-matBackend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProdcutHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProdcutHandler {
	return &ProdcutHandler{
		service: service,
	}
}

func (h *ProdcutHandler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}

	product, err := h.service.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProdcutHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch products",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}
