package mysql

import (
	"E-matBackend/internal/models"
	"database/sql"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	query := `SELECT id,name,Price,Image,category,isNew FROM products WHERE id = ? `
	row := r.db.QueryRow(query, id)

	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Image, &product.Category, &product.IsNew)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	query := `SELECT * FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Image, &p.Category, &p.IsNew)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
