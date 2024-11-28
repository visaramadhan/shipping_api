package cmd

import (
	"github.com/visaramadhan/shipping_api.git/api/order_shiping"
	"github.com/visaramadhan/shipping_api.git/api/shiping"
	"gorm.io/gorm"
	"log"
	"time"
)

// SeedOrderShipping untuk mengisi data OrderShipping ke database
func SeedOrderShipping(db *gorm.DB) {
	// Data dummy untuk OrderShipping
	orderShippings := []order_shiping.OrderShipping{
		{
			ID:              "order-001",
			EcommerceID:     "ecommerce-001",
			ShippingID:      "shipping-001",
			Shipping:        shiping.Shiping{ID: "shipping-001", Name: "Courier Express"},
			OriginLongitude: 106.7952,
			DestinationLat:  -6.2088,
			DestinationLong: 106.8456,
			TotalPayment:    150000,
			OrderDate:       time.Now(),
		},
		{
			ID:              "order-002",
			EcommerceID:     "ecommerce-002",
			ShippingID:      "shipping-002",
			Shipping:        shiping.Shiping{ID: "shipping-002", Name: "Fast Delivery"},
			OriginLongitude: 106.8272,
			DestinationLat:  -6.1745,
			DestinationLong: 106.8221,
			TotalPayment:    200000,
			OrderDate:       time.Now(),
		},
		// Tambahkan lebih banyak data dummy sesuai kebutuhan
	}

	// Menyisipkan data ke dalam database jika belum ada
	for _, order := range orderShippings {
		// Cek apakah data sudah ada di database berdasarkan ID
		var existingOrder order_shiping.OrderShipping
		err := db.Where("id = ?", order.ID).First(&existingOrder).Error
		if err == nil {
			// Jika sudah ada, lewati dan lanjutkan
			log.Printf("Order with ID %s already exists, skipping...\n", order.ID)
			continue
		}

		// Jika belum ada, insert data baru
		err = db.Create(&order).Error
		if err != nil {
			log.Printf("Error seeding OrderShipping with ID %s: %v\n", order.ID, err)
		} else {
			log.Printf("Successfully seeded OrderShipping with ID %s\n", order.ID)
		}
	}
}
