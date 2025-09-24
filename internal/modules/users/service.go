package users

import (
	"github.com/arthu/shop-api-go/internal/models"
	"github.com/arthu/shop-api-go/internal/utils"
)

// Service exposes business operations for users domain.
type Service interface {
	Create(u *models.User) (int64, error)
	List() ([]models.User, error)
	GetByID(id int64) (*models.User, error)
	Update(id int64, u *models.User) (bool, error)
	Delete(id int64) (bool, error)
}

type service struct{ repo Repository }

func NewService(r Repository) Service { return &service{repo: r} }

func (s *service) Create(u *models.User) (int64, error) {
	if u.Password != "" {
		h, err := utils.HashPassword(u.Password)
		if err != nil { return 0, err }
		u.Password = h
	}
	return s.repo.CreateUser(u)
}

func (s *service) List() ([]models.User, error)              { return s.repo.ListUsers() }
func (s *service) GetByID(id int64) (*models.User, error)    { return s.repo.GetUserByID(id) }
func (s *service) Update(id int64, u *models.User) (bool, error) { return s.repo.UpdateUser(id, u) }
func (s *service) Delete(id int64) (bool, error)            { return s.repo.DeleteUser(id) }
