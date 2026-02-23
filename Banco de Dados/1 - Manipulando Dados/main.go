package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewPrduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "golang:golang@/goexpert?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewPrduct("Adidas Campus 00s", 799.00)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Price = 499.00
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("O preço do produto: %s é de R$ %v", p.Name, p.Price))

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}

	for _, p := range products {
		fmt.Println(p)
	}

	fmt.Println("---------------------------------------")
	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	products, err = selectAllProducts(db)
	if err != nil {
		panic(err)
	}

	for _, p := range products {
		fmt.Println(p)
	}

	err = deleteAllProduct(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("---------------------------------------")
	products, err = selectAllProducts(db)
	if err != nil {
		panic(err)
	}

	fmt.Println(products)
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select * from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select * from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func deleteAllProduct(db *sql.DB) error {
	_, err := db.Exec("delete from products")
	if err != nil {
		return err
	}
	return nil
}
