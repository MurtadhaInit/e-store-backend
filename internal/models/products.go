package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float32
	Stock       int
	Added       time.Time
}

type ProductModel struct {
	DB *sql.DB
}

func (m *ProductModel) Insert(title string, description string, price float32, stock int) (int, error) {
	stmt := `INSERT INTO products (title, description, price, stock_quantity, date_added)
					VALUES (?, ?, ?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(stmt, title, description, price, stock)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
