package models

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Image    string  `json:"image"`
	Category string  `json:"category"`
	IsNew    bool    `json:"isNew"`
}

type ProductSlice []Product
