package shiping

import (
	"errors"
	"github.com/visaramadhan/shipping_api.git/api/destination"
	"gorm.io/gorm"
)

type ShipingRepository interface {
	CreateNewShiping(payload *Shiping) (*Shiping, error)
	GetAllShipings() ([]Shiping, error) // Tidak perlu pointer ke slice
	GetShipingById(id string) (*Shiping, error)
	CalculateCost(payload destination.RequestDestination) (float64, error)   // Mengembalikan nilai langsung, bukan pointer
	GetDestination(payload destination.RequestDestination) (*float64, error) // Menambahkan method GetDestination
}

type shipingRepository struct {
	db *gorm.DB
}

func NewShipingRepository(db *gorm.DB) ShipingRepository {
	return &shipingRepository{db: db}
}

// CreateNewShiping menyimpan data pengiriman baru ke dalam database
func (s *shipingRepository) CreateNewShiping(payload *Shiping) (*Shiping, error) {
	result := s.db.Create(&payload)
	if result.Error != nil {
		return nil, result.Error
	}
	return payload, nil
}

// GetAllShipings mengambil semua data pengiriman dari database
func (s *shipingRepository) GetAllShipings() ([]Shiping, error) {
	var shipings []Shiping
	result := s.db.Find(&shipings)
	if result.Error != nil {
		return nil, result.Error
	}
	return shipings, nil
}

// GetShipingById mengambil data pengiriman berdasarkan ID
func (s *shipingRepository) GetShipingById(id string) (*Shiping, error) {
	var shiping Shiping
	result := s.db.First(&shiping, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &shiping, nil
}

// CalculateCost menghitung biaya pengiriman berdasarkan data yang diberikan
func (s *shipingRepository) CalculateCost(payload destination.RequestDestination) (float64, error) {
	// Mendapatkan jarak dari fungsi repository destination
	distance, err := s.GetDestination(payload)
	if err != nil {
		return 0, err
	}

	// Periksa apakah jarak valid
	if distance == nil || *distance <= 0 {
		return 0, errors.New("invalid distance value")
	}

	// Kalkulasi ongkos berdasarkan jumlah barang dan jarak
	var costPerKm float64
	if payload.Qty <= 2 {
		costPerKm = 2000
	} else {
		costPerKm = 4000
	}

	cost := *distance * costPerKm
	return cost, nil
}

// GetDestination akan mengakses fungsi atau service untuk mendapatkan jarak (mungkin menggunakan API atau data lainnya)
func (s *shipingRepository) GetDestination(payload destination.RequestDestination) (*float64, error) {
	var distance float64
	distance = 100.0 // Anggap jarak sementara adalah 100 km

	return &distance, nil
}
