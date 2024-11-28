package order_shiping

import (
	"errors"
	"gorm.io/gorm"
)

// OrderRepository interface untuk operasi CRUD pada OrderShipping
type OrderRepository interface {
	Create(payload *OrderShipping) (*OrderShipping, error) // Menggunakan pointer ke OrderShipping
	Get(id string) (*OrderShipping, error)
	Update(id string, payload *OrderShipping) (*OrderShipping, error) // Menggunakan pointer ke OrderShipping
}

// orderRepository adalah implementasi dari OrderRepository
type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository membuat instance baru dari orderRepository
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

// Create menyimpan OrderShipping baru ke dalam database
func (os *orderRepository) Create(payload *OrderShipping) (*OrderShipping, error) {
	// Menetapkan EcommerceID secara hardcode
	payload.EcommerceID = "Ecommerce123" // Hardcode Ecommerce ID

	// Pastikan nilai EcommerceID sudah terisi sebelum membuat entri baru
	if payload.EcommerceID == "" {
		return nil, errors.New("EcommerceID cannot be empty")
	}

	// Simpan ke database
	err := os.db.Create(payload).Error
	return payload, err
}

// Get mengambil OrderShipping berdasarkan ID
func (os *orderRepository) Get(id string) (*OrderShipping, error) {
	var payload OrderShipping
	// Ambil data OrderShipping berdasarkan ID
	err := os.db.Where("id = ?", id).First(&payload).Error
	return &payload, err
}

// Update memperbarui OrderShipping berdasarkan ID
func (os *orderRepository) Update(id string, payload *OrderShipping) (*OrderShipping, error) {
	// Pastikan EcommerceID diupdate dengan nilai yang benar
	if payload.EcommerceID == "" {
		return nil, errors.New("EcommerceID cannot be empty")
	}

	// Perbarui data OrderShipping
	err := os.db.Where("id = ?", id).Updates(payload).Error
	return payload, err
}
