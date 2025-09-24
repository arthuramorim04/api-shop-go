package users

import (
	"github.com/arthu/shop-api-go/internal/models"
	"github.com/arthu/shop-api-go/internal/repo"
)

// Repository abstracts persistence for users domain.
type Repository interface {
	CreateUser(u *models.User) (int64, error)
	GetUserByID(id int64) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	ListUsers() ([]models.User, error)
	UpdateUser(id int64, u *models.User) (bool, error)
	DeleteUser(id int64) (bool, error)
}

// repoAdapter adapts existing functions in internal/repo.
type repoAdapter struct{}

func NewRepository() Repository { return &repoAdapter{} }

func (repoAdapter) CreateUser(u *models.User) (int64, error)              { return repo.CreateUser(u) }
func (repoAdapter) GetUserByID(id int64) (*models.User, error)            { return repo.GetUserByID(id) }
func (repoAdapter) GetUserByEmail(email string) (*models.User, error)     { return repo.GetUserByEmail(email) }
func (repoAdapter) ListUsers() ([]models.User, error)                     { return repo.ListUsers() }
func (repoAdapter) UpdateUser(id int64, u *models.User) (bool, error)     { return repo.UpdateUser(id, u) }
func (repoAdapter) DeleteUser(id int64) (bool, error)                     { return repo.DeleteUser(id) }
