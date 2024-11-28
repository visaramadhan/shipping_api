package shiping

import (
	"errors"
	"fmt"
	"github.com/visaramadhan/shipping_api.git/api/destination"
)

type ShipingService interface {
	CreateNewShiping(payload *Shiping) (*Shiping, error)
	GetAllShipings() ([]Shiping, error) // Tidak perlu pointer ke slice
	GetShipingById(id string) (*Shiping, error)
	CalculateCost(payload destination.RequestDestination) (float64, error) // Mengembalikan nilai langsung, bukan pointer
}

type shipingService struct {
	repo ShipingRepository
}

func NewShipingService(repo ShipingRepository) ShipingService {
	return &shipingService{repo: repo}
}

// CreateNewShiping membuat pengiriman baru
func (s *shipingService) CreateNewShiping(payload *Shiping) (*Shiping, error) {
	// Validasi input
	if payload.Name == "" {
		return nil, errors.New("name is required")
	}
	// Pastikan ada detail alamat atau data lain jika diperlukan
	if payload.Address == "" {
		return nil, errors.New("address is required")
	}

	// Menyimpan data pengiriman baru
	return s.repo.CreateNewShiping(payload)
}

// GetAllShipings mengambil seluruh data pengiriman
func (s *shipingService) GetAllShipings() ([]Shiping, error) { // Mengembalikan slice, bukan pointer ke slice
	shipings, err := s.repo.GetAllShipings()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve shipings: %w", err)
	}
	return shipings, nil
}

// GetShipingById mengambil data pengiriman berdasarkan ID
func (s *shipingService) GetShipingById(id string) (*Shiping, error) {
	data, err := s.repo.GetShipingById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve shiping by id: %w", err)
	}
	return data, nil
}
func (s *shipingService) CalculateCost(payload destination.RequestDestination) (float64, error) {
	// Mendapatkan jarak dari repository
	distance, err := s.repo.GetDestination(payload)
	if err != nil {
		return 0, fmt.Errorf("failed to get destination: %w", err)
	}

	// Pastikan jarak valid
	if distance == nil || *distance <= 0 {
		return 0, errors.New("invalid distance value")
	}

	// Pastikan jumlah barang ada dalam payload
	if payload.Qty <= 0 {
		return 0, errors.New("invalid quantity value")
	}

	// Kalkulasi ongkos pengiriman berdasarkan jumlah barang dan jarak
	var costPerKm float64
	if payload.Qty <= 2 {
		costPerKm = 2000 // Ongkos per km untuk 2 barang atau kurang
	} else {
		costPerKm = 4000 // Ongkos per km untuk lebih dari 2 barang
	}

	// Hitung total biaya
	cost := *distance * costPerKm
	return cost, nil
}
