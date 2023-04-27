package product2

import (
	"errors"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

var (
	ErrorNotFound    = errors.New("user not found in the given repository")
	ErrAlreadyExists = errors.New("user already exists in the given repository")
)

type Repository2 interface {
	GetByID(id int) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(id int, product *domain.Product) error
	Delete(id int) error
}
