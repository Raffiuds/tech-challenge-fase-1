package repositories

import (
	"tech-challenge-fase-1/internal/core/domain/entities"
	valueobjects "tech-challenge-fase-1/internal/core/domain/value_objects"
)

type CustomerRepositoryInterface interface {
	GetCustomerByCPF(cpf *valueobjects.CPF) (*entities.Customer, error)
	GetCustomerByEmail(email *valueobjects.Email) (*entities.Customer, error)
	Insert(*entities.Customer) error
}
