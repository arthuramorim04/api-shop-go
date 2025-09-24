package products

import (
	"github.com/arthu/shop-api-go/internal/models"
	"github.com/arthu/shop-api-go/internal/repo"
)

// Repository abstracts persistence for products domain.
type Repository interface {
	CreateProduct(p *models.Product) (int64, error)
	ListProducts() ([]models.Product, error)
	GetProductByID(id int64) (*models.Product, error)
	UpdateProduct(id int64, p *models.Product) (bool, error)
	DeleteProduct(id int64) (bool, error)
}

type repoAdapter struct{}

func NewRepository() Repository { return &repoAdapter{} }

func (repoAdapter) CreateProduct(p *models.Product) (int64, error)     { return repo.CreateProduct(p) }
func (repoAdapter) ListProducts() ([]models.Product, error)            { return repo.ListProducts() }
func (repoAdapter) GetProductByID(id int64) (*models.Product, error)   { return repo.GetProductByID(id) }
func (repoAdapter) UpdateProduct(id int64, p *models.Product) (bool, error) { return repo.UpdateProduct(id, p) }
func (repoAdapter) DeleteProduct(id int64) (bool, error)               { return repo.DeleteProduct(id) }
