package history_delivery

import (
	"gorm.io/gorm"
)

type HistoryRepo interface {
	DeliverHistory(id History) error
	GetHistory(id History) ([]History, error)
	Tracking(id History) ([]History, error)
	UpdateStatus(payload History) (History, error)
}

type historyRepo struct {
	db *gorm.DB
}

// NewHistoryRepo creates a new instance of HistoryRepo
func NewHistoryRepo(db *gorm.DB) HistoryRepo {
	return &historyRepo{db: db}
}

// DeliverHistory deletes a history record based on the given History ID
func (hr *historyRepo) DeliverHistory(id History) error {
	// Perform deletion based on History ID
	db := hr.db.Where("id = ?", id.ID).Delete(&id)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// GetHistory retrieves history records based on the given History ID
func (hr *historyRepo) GetHistory(id History) ([]History, error) {
	var histories []History
	// Fetch the histories with the related OrderShipping using Preload
	db := hr.db.Where("id = ?", id.ID).Preload("Order").Find(&histories)
	if db.Error != nil {
		return nil, db.Error
	}
	return histories, nil
}

// Tracking retrieves history records based on the given Tracking ID
func (hr *historyRepo) Tracking(id History) ([]History, error) {
	var histories []History
	// Fetch the histories with the related OrderShipping using Preload
	db := hr.db.Where("id = ?", id.ID).Preload("Order").Find(&histories)
	if db.Error != nil {
		return nil, db.Error
	}
	return histories, nil
}

// UpdateStatus updates the status of a specific history record
func (hr *historyRepo) UpdateStatus(payload History) (History, error) {
	// Update the status of a history record based on the given ID
	db := hr.db.Where("id = ?", payload.ID).Save(&payload)
	if db.Error != nil {
		return History{}, db.Error
	}
	return payload, nil
}
