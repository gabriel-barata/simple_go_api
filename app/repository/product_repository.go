package repository

import (
	"database/sql"
	"fmt"
	"log"
	"simple-go-api/app/models"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
	query := "SELECT id, product_name, price FROM products"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []models.Product{}, err
	}

	var productList []models.Product
	var productObject models.Product

	for rows.Next() {
		err = rows.Scan(
			&productObject.ID,
			&productObject.Name,
			&productObject.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []models.Product{}, nil
		}
	}

	productList = append(productList, productObject)

	rows.Close()

	return productList, nil

}

func (pr *ProductRepository) InsertProduct(product models.Product) (int, error) {
	var id int

	query, err := pr.connection.Prepare(
		"INSERT INTO products (product_name, price) VALUES ($1, $2) RETURNING id",
	)

	if err != nil {
		log.Fatalf("Unable to prepare insert query: $%v", err)
		return id, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id) // pointing to id(var) memory address
	if err != nil {
		log.Fatalf("Unable to insert values on table: %v", err)
		return id, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductById(productId int) (*models.Product, error) {
	var product models.Product

	query, err := pr.connection.Prepare(
		"SELECT * FROM products WHERE id = $1",
	)
	if err != nil {
		log.Fatalf("Unable to prepare query: %v", err)
		return nil, err
	}

	err = query.QueryRow(productId).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Fatalf("Unable to parse query result into models.Product: %v", err)
		return nil, err
	}
	query.Close()

	return &product, nil
}
