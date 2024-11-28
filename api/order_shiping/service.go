package order_shiping

import (
	"errors"
)

// OrderService adalah interface yang mendefinisikan metode-metode layanan order.
type OrderService interface {
	Create(payload *OrderShipping) (*OrderShipping, error)
	Get(id string) (*OrderShipping, error)
	Update(id string, payload *OrderShipping) (*OrderShipping, error)
}

// orderService adalah implementasi OrderService yang menggunakan OrderRepository.
type orderService struct {
	repo OrderRepository
}

// NewOrderService adalah konstruktor untuk OrderService.
func NewOrderService(repo OrderRepository) OrderService {
	return &orderService{repo: repo}
}

// Create membuat order baru.
func (s *orderService) Create(payload *OrderShipping) (*OrderShipping, error) {
	if payload == nil {
		return nil, errors.New("payload cannot be nil")
	}

	// Memanggil metode Create dari repository
	return s.repo.Create(payload)
}

// Get mengambil order berdasarkan ID.
func (s *orderService) Get(id string) (*OrderShipping, error) {
	if id == "" {
		return nil, errors.New("ID cannot be empty")
	}

	// Memanggil metode Get dari repository
	return s.repo.Get(id)
}

// Update memperbarui order berdasarkan ID dan payload.
func (s *orderService) Update(id string, payload *OrderShipping) (*OrderShipping, error) {
	if id == "" {
		return nil, errors.New("ID cannot be empty")
	}
	if payload == nil {
		return nil, errors.New("payload cannot be nil")
	}

	// Memanggil metode Update dari repository
	return s.repo.Update(id, payload)
}
