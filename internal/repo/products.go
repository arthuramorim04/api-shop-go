package repo

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/arthu/shop-api-go/internal/database"
	"github.com/arthu/shop-api-go/internal/models"
)

func CreateProduct(p *models.Product) (int64, error) {
	s := `INSERT INTO products (name, description, price, quantity, imageUrl) VALUES (?,?,?,?,?)`
	res, err := database.DB().Exec(s, p.Name, p.Description, p.Price, p.Quantity, p.ImageURL)
	if err != nil { return 0, err }
	return res.LastInsertId()
}

func ListProducts() ([]models.Product, error) {
	rows, err := database.DB().Query(`SELECT id, name, description, price, quantity, imageUrl FROM products`)
	if err != nil { return nil, err }
	defer rows.Close()
	var out []models.Product
	for rows.Next() {
		p := models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.ImageURL); err != nil { return nil, err }
		out = append(out, p)
	}
	return out, nil
}

func GetProductByID(id int64) (*models.Product, error) {
	row := database.DB().QueryRow(`SELECT id, name, description, price, quantity, imageUrl FROM products WHERE id = ?`, id)
	p := models.Product{}
	if err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.ImageURL); err != nil {
		if err == sql.ErrNoRows { return nil, nil }
		return nil, err
	}
	return &p, nil
}

func UpdateProduct(id int64, p *models.Product) (bool, error) {
	s := `UPDATE products SET name=?, description=?, price=?, quantity=?, imageUrl=? WHERE id=?`
	res, err := database.DB().Exec(s, p.Name, p.Description, p.Price, p.Quantity, p.ImageURL, id)
	if err != nil { return false, err }
	affected, _ := res.RowsAffected()
	return affected > 0, nil
}

func DeleteProduct(id int64) (bool, error) {
	res, err := database.DB().Exec(`DELETE FROM products WHERE id=?`, id)
	if err != nil { return false, err }
	affected, _ := res.RowsAffected()
	return affected > 0, nil
}

func GetProductsByIDs(ids []int64) ([]models.Product, error) {
	if len(ids) == 0 { return []models.Product{}, nil }
	// Prepare IN clause
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids { placeholders[i] = "?"; args[i] = id }
	s := fmt.Sprintf("SELECT id, name, description, price, quantity, imageUrl FROM products WHERE id IN (%s)", 
		strings.Join(placeholders, ","))
	rows, err := database.DB().Query(s, args...)
	if err != nil { return nil, err }
	defer rows.Close()
	var out []models.Product
	for rows.Next() {
		p := models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.ImageURL); err != nil { return nil, err }
		out = append(out, p)
	}
	return out, nil
}

func GetProductsForOrder(orderID string) ([]models.Product, error) {
	s := `SELECT p.id, p.name, p.description, p.price, p.quantity, p.imageUrl
		FROM products p JOIN order_products op ON p.id = op.product_id WHERE op.order_id = ?`
	rows, err := database.DB().Query(s, orderID)
	if err != nil { return nil, err }
	defer rows.Close()
	var out []models.Product
	for rows.Next() {
		p := models.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.ImageURL); err != nil { return nil, err }
		out = append(out, p)
	}
	return out, nil
}
