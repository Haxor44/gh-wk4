package services

import (
	"E-matBackend/internal/models"
	"E-matBackend/internal/repositories/mysql"
	"E-matBackend/internal/repositories/redis"
	"strconv"
	"time"
)

type ProductService struct {
	productRepo *mysql.ProductRepository
	cacheRepo   *redis.CacheRepository
}

func NewProductService(pr *mysql.ProductRepository, cr *redis.CacheRepository) *ProductService {
	return &ProductService{
		productRepo: pr,
		cacheRepo:   cr,
	}
}

func (s *ProductService) GetProduct(id int) (*models.Product, error) {
	cacheKey := "product:" + strconv.Itoa(id)

	// Try cache first
	var cachedProduct models.Product
	err := s.cacheRepo.Get(cacheKey, &cachedProduct)
	if err == nil {
		return &cachedProduct, nil
	}

	// Cache miss -> fetch from the DB
	product, err := s.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	// update cache
	s.cacheRepo.Set(cacheKey, product, 10*time.Minute)

	return product, nil

}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	cacheKey := "products:all"

	// Try cache first
	var cachedProducts []models.Product
	err := s.cacheRepo.GetSlice(cacheKey, &cachedProducts)
	if err == nil {
		return cachedProducts, nil
	}

	// Cache miss - fetch from DB
	products, err := s.productRepo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	// Update cache
	s.cacheRepo.SetSlice(cacheKey, products, 10*time.Minute)

	return products, nil
}
