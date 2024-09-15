package service

import (
	"github.com/yaken-org/hakushi/internal/database"
	"github.com/yaken-org/hakushi/internal/model"
)

func FindProductByID(id int64) (*model.Product, error) {
	db := database.New()
	product := new(model.Product)
	if err := model.QueryRow(db.DB, product, "SELECT * FROM product WHERE id = ?", id); err != nil {
		return nil, err
	}
	return product, nil
}

func FindAllProducts() ([]*model.Product, error) {
	db := database.New()
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		return nil, err
	}

	products := make([]*model.Product, 0)
	for rows.Next() {
		product := new(model.Product)
		if err := product.FromRow(rows); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func CreateProduct(name string, link string) (*model.Product, error) {
	db := database.New()
	res, err := db.Exec(`
		INSERT INTO product (name, link) VALUES(?, ?)
	`, name, link)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return FindProductByID(id)
}

func DeleteProduct(id int64) error {
	db := database.New()
	_, err := db.Exec("DELETE FROM product WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(id int64, name string, link string) (*model.Product, error) {
	db := database.New()
	_, err := db.Exec("UPDATE product SET name = ?, link = ? WHERE id = ?", name, link, id)
	if err != nil {
		return nil, err
	}
	return FindProductByID(id)
}
