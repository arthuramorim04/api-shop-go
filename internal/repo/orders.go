package repo

import (
	"database/sql"

	"github.com/arthu/shop-api-go/internal/database"
	"github.com/arthu/shop-api-go/internal/models"
)

func InsertOrder(o *models.Order) error {
	s := `INSERT INTO orders (id, user_id, order_date, status, payment_url) VALUES (?,?,?,?,?)`
	_, err := database.DB().Exec(s, o.ID, o.UserID, o.OrderDate, o.Status, o.PaymentURL)
	return err
}

func InsertOrderProducts(orderID string, items []struct{ ProductID int64; Quantity int }) error {
	s := `INSERT INTO order_products (order_id, product_id, quantity) VALUES (?,?,?)`
	for _, it := range items {
		if _, err := database.DB().Exec(s, orderID, it.ProductID, it.Quantity); err != nil { return err }
	}
	return nil
}

func ListOrders() ([]models.Order, error) {
	rows, err := database.DB().Query(`SELECT id, user_id, order_date, status, payment_url FROM orders`)
	if err != nil { return nil, err }
	defer rows.Close()
	var out []models.Order
	for rows.Next() {
		o := models.Order{}
		if err := rows.Scan(&o.ID, &o.UserID, &o.OrderDate, &o.Status, &o.PaymentURL); err != nil { return nil, err }
		out = append(out, o)
	}
	return out, nil
}

func GetOrderByID(id string) (*models.Order, error) {
	row := database.DB().QueryRow(`SELECT id, user_id, order_date, status, payment_url FROM orders WHERE id = ?`, id)
	o := models.Order{}
	if err := row.Scan(&o.ID, &o.UserID, &o.OrderDate, &o.Status, &o.PaymentURL); err != nil {
		if err == sql.ErrNoRows { return nil, nil }
		return nil, err
	}
	return &o, nil
}

func GetOrdersByUserID(userID int64) ([]models.Order, error) {
	rows, err := database.DB().Query(`SELECT id, user_id, order_date, status, payment_url FROM orders WHERE user_id = ?`, userID)
	if err != nil { return nil, err }
	defer rows.Close()
	var out []models.Order
	for rows.Next() {
		o := models.Order{}
		if err := rows.Scan(&o.ID, &o.UserID, &o.OrderDate, &o.Status, &o.PaymentURL); err != nil { return nil, err }
		out = append(out, o)
	}
	return out, nil
}

func GetOrdersByProductID(productID int64) ([]models.Order, error) {
	s := `SELECT o.id, o.user_id, o.order_date, o.status, o.payment_url 
	FROM orders o JOIN order_products op ON o.id = op.order_id WHERE op.product_id = ?`
	rows, err := database.DB().Query(s, productID)
	if err != nil { return nil, err }
	defer rows.Close()
	var out []models.Order
	for rows.Next() {
		o := models.Order{}
		if err := rows.Scan(&o.ID, &o.UserID, &o.OrderDate, &o.Status, &o.PaymentURL); err != nil { return nil, err }
		out = append(out, o)
	}
	return out, nil
}

func UpdateOrderStatus(orderID string, status models.OrderStatus) (bool, error) {
	res, err := database.DB().Exec(`UPDATE orders SET status = ? WHERE id = ?`, status, orderID)
	if err != nil { return false, err }
	affected, _ := res.RowsAffected()
	return affected > 0, nil
}
