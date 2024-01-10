package db_test

import (
	"arquitetura-hexagonal/adapters/db"
	"arquitetura-hexagonal/application"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {

	sqlCreateTable := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
		);`

	stmt, err := db.Prepare(sqlCreateTable)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("ar1te","Smartphone Galaxy S10", 2000.50,"ENABLED")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("ar1te")

	require.Nil(t, err)
	require.Equal(t, "Smartphone Galaxy S10", product.GetName())
	require.Equal(t, 2000.50, product.GetPrice())
	require.Equal(t, "ENABLED", product.GetStatus())
}

func TestProduct_Save(t *testing.T) {

	t.Run("Save new product", func(t *testing.T) {

		setUp()
		defer Db.Close()
		productDb := db.NewProductDb(Db)

		newProduct := application.NewProduct()
		newProduct.Name = "Iphone 15 Pro Max"
		newProduct.Price = 8000.0
		newProduct.Status = application.ENABLED

		_, err := productDb.Save(newProduct)

		require.Nil(t, err)

		productSaved, err := productDb.Get(newProduct.GetID())

		require.Nil(t, err)
		require.NotNil(t, productSaved.GetID())
		require.Equal(t, newProduct.GetName(), productSaved.GetName())
		require.Equal(t, newProduct.GetPrice(), productSaved.GetPrice())
		require.Equal(t, newProduct.GetStatus(), productSaved.GetStatus())

	})

	t.Run("Update the product", func(t *testing.T) {

		setUp()
		defer Db.Close()
		productDb := db.NewProductDb(Db)

		productToUpdate := application.NewProduct()
		productToUpdate.ID = "ar1te"
		productToUpdate.Name = "Iphone 14"
		productToUpdate.Price = 4000.0

		_, err := productDb.Save(productToUpdate)

		require.Nil(t, err)

		productUpdated, err := productDb.Get("ar1te")

		require.Nil(t, err)
		require.Equal(t, productToUpdate.GetID(), productUpdated.GetID())
		require.Equal(t, productToUpdate.GetName(), productUpdated.GetName())
		require.Equal(t, productToUpdate.GetPrice(), productUpdated.GetPrice())
		require.Equal(t, productToUpdate.GetStatus(), productUpdated.GetStatus())

	})
}
