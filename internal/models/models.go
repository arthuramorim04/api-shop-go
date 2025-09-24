package models

import "time"

type Role string

const (
	RoleClient  Role = "client"
	RoleSupport Role = "suport" // mantém o mesmo nome usado no SQL original
	RoleAdmin   Role = "admin"
)

type User struct {
	ID          int64  `json:"id" db:"id"`
	Username    string `json:"username" db:"username"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password,omitempty" db:"password"`
	FirstName   string `json:"firstName" db:"firstName"`
	LastName    string `json:"lastName" db:"lastName"`
	Address     string `json:"address" db:"address"`
	PhoneNumber string `json:"phoneNumber" db:"phoneNumber"`
	Role        Role   `json:"role" db:"role"`
}

type UserDTO struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Role        Role   `json:"role"`
	Token       string `json:"token"`
}

type Product struct {
	ID          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Quantity    int     `json:"quantity" db:"quantity"`
	ImageURL    string  `json:"imageUrl" db:"imageUrl"`
}

type OrderStatus string

const (
	OrderPending   OrderStatus = "pending"
	OrderShipped   OrderStatus = "shipped"
	OrderDelivered OrderStatus = "delivered"
	OrderError     OrderStatus = "error"
	OrderApproved  OrderStatus = "approved"
)

type OrderProduct struct {
	ID        int64 `json:"id"`
	OrderID   string `json:"order_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type Order struct {
	ID         string       `json:"id" db:"id"`
	UserID     int64        `json:"userId" db:"user_id"`
	OrderDate  time.Time    `json:"orderDate" db:"order_date"`
	PaymentURL string       `json:"payment_url" db:"payment_url"`
	Status     OrderStatus  `json:"status" db:"status"`
	Products   []Product    `json:"products"`
}
