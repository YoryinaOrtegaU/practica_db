package product

import (
	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/bootcamp-go/consignas-go-db.git/internal/product2"
)

type Service interface {
	// GetByID busca un producto por su id
	GetByID(id int) (domain.Product, error)
	// Create agrega un nuevo producto
	Create(p domain.Product) (domain.Product, error)
	// Delete elimina un producto
	Delete(id int) error
	// Update actualiza un producto
	Update(id int, p domain.Product) (domain.Product, error)
}

/* type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
} */

type service struct {
	r product2.Repository2
}

func NewService(r product2.Repository2) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Product, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return *p, nil
}

func (s *service) Create(p domain.Product) (domain.Product, error) {
	err := s.r.Create(&p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}
func (s *service) Update(id int, u domain.Product) (domain.Product, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	if u.Name != "" {
		p.Name = u.Name
	}
	if u.CodeValue != "" {
		p.CodeValue = u.CodeValue
	}
	if u.Expiration != "" {
		p.Expiration = u.Expiration
	}
	if u.Quantity > 0 {
		p.Quantity = u.Quantity
	}
	if u.Price > 0 {
		p.Price = u.Price
	}
	err = s.r.Update(id, p)
	if err != nil {
		return domain.Product{}, err
	}
	return *p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
