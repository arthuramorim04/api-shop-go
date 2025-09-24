package objectmothers

import (
	"time"

	"github.com/arthu/shop-api-go/internal/models"
)

func User() models.User {
	return models.User{
		ID:          1,
		Username:    "john",
		Email:       "john@example.com",
		Password:    "$2a$10$abcdefghijklmnopqrstuv", // fake hash
		FirstName:   "John",
		LastName:    "Doe",
		Address:     "123 Main St",
		PhoneNumber: "+55 11 99999-9999",
		Role:        models.RoleAdmin,
	}
}

func Product() models.Product {
	return models.Product{
		ID:          10,
		Name:        "Widget",
		Description: "A useful widget",
		Price:       99.90,
		Quantity:    5,
		ImageURL:    "https://example.com/widget.png",
	}
}

func Order() models.Order {
	return models.Order{
		ID:         "order-123",
		UserID:     1,
		OrderDate:  time.Now(),
		PaymentURL: "https://pay.example.com/abc",
		Status:     models.OrderPending,
	}
}
