package infrastructure

import (
	"database/sql"
	"errors"
	"github.com/martinezgomez34/abarrote/src/core"
	"github.com/martinezgomez34/abarrote/src/products/domain"
)

type MySql struct {
	db *core.ConnMySQL
}

func NewSQL(db *core.ConnMySQL) *MySql {

	return &MySql{db: db}
}

func (r *MySql) Save(product *domain.Product) error {
	query := `INSERT INTO products (name, price, amount, brand, description) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.DB.Exec(query, product.Name, product.Price, product.Amount, product.Brand, product.Description)
	return err
}

func (r *MySql) GetAll() ([]*domain.Product, error) {
	query := `SELECT id, name, price, amount, brand, description FROM products`
	rows, err := r.db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var p domain.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Amount, &p.Brand, &p.Description)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}

func (r *MySql) GetByID(id int32) (*domain.Product, error) {
	query := `SELECT id, name, price, amount, brand, description FROM products WHERE id = ?`
	row := r.db.DB.QueryRow(query, id)

	var product domain.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Amount, &product.Brand, &product.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (r *MySql) Update(product *domain.Product) error {
	query := `UPDATE products SET name = ?, price = ?, amount = ?, brand = ?, description = ? WHERE id = ?`
	_, err := r.db.DB.Exec(query, product.Name, product.Price, product.Amount, product.Brand, product.Description, product.ID)
	return err
}

func (r *MySql) Delete(id int32) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := r.db.DB.Exec(query, id)
	return err
}
