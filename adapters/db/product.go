package db

import (
	"arquitetura-hexagonal/application"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {

	var product application.Product

	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {

	var productReturned application.ProductInterface

	isNewProduct, err := p.isNewProduct(product)

	if err != nil {
		return nil, err
	}

	if isNewProduct {

		productReturned, err = p.create(product)

		if err != nil {
			return nil, err
		}

		return productReturned, nil

	}

	productReturned, err = p.update(product)

	if err != nil {
		return nil, err
	}

	return productReturned, nil

}

func (p *ProductDb) isNewProduct(product application.ProductInterface) (bool, error) {

	var rowsReturned int

	stmt, err := p.db.Prepare("SELECT COUNT(*) FROM products WHERE id=?")

	if err != nil {
		return false, err
	}

	err = stmt.QueryRow(product.GetID()).Scan(&rowsReturned)

	if err != nil {
		return false, err
	}

	return rowsReturned == 0, nil
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {

	createdProductSql := "INSERT INTO products(id, name, price, status) VALUES (?, ?, ?, ?)"

	stmt, err := p.db.Prepare(createdProductSql)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {

	updateProductSql := "UPDATE products SET price = ?, name = ?, status = ? WHERE id = ?"

	_, err := p.db.Exec(
		updateProductSql,
		product.GetPrice(),
		product.GetName(),
		product.GetStatus(),
		product.GetID(),
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
