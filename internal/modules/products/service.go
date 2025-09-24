package products

import (
	"github.com/arthu/shop-api-go/internal/models"
)

// Service exposes business operations for products domain.
type Service interface {
	Create(p *models.Product) (int64, error)
	List() ([]models.Product, error)
	GetByID(id int64) (*models.Product, error)
	Update(id int64, p *models.Product) (bool, error)
	Delete(id int64) (bool, error)
}

type service struct{ repo Repository }

func NewService(r Repository) Service { return &service{repo: r} }

func (s *service) Create(p *models.Product) (int64, error)           { return s.repo.CreateProduct(p) }
func (s *service) List() ([]models.Product, error)                    { return s.repo.ListProducts() }
func (s *service) GetByID(id int64) (*models.Product, error)          { return s.repo.GetProductByID(id) }
func (s *service) Update(id int64, p *models.Product) (bool, error)   { return s.repo.UpdateProduct(id, p) }
func (s *service) Delete(id int64) (bool, error)                      { return s.repo.DeleteProduct(id) }
