package shiping

import "gorm.io/gorm"

type ShipingRepository interface {
	Create(payload Shiping) (Shiping, error)
}

type shipingRepository struct {
	db *gorm.DB
}

func NewShipingRepository(db *gorm.DB) shipingRepository {
	return shipingRepository{db: db}
}

func (s shipingRepository) Create(payload Shiping) (Shiping, error) {
	result := s.db.Create(&payload)
    if result.Error != nil {
        return Shiping{}, result.Error
    }
    return payload, nil
}