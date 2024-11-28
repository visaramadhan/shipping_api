package main

import (
	"fmt"
	"log"
	"time"

	"github.com/visaramadhan/shipping_api.git/api/order_shiping"
	"github.com/visaramadhan/shipping_api.git/api/shiping"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Inisialisasi DB GORM (SQLite)
	db, err := gorm.Open(sqlite.Open("shipping.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Auto migrate untuk memastikan tabel sudah ada
	db.AutoMigrate(&order_shiping.OrderShipping{})

	// Inisialisasi OrderRepository
	repo := order_shiping.NewOrderRepository(db)

	// Membuat OrderShipping baru
	order := &order_shiping.OrderShipping{
		ID:              "ORDER-1-SHP-281124",
		ShippingID:      "shipping001",
		OriginLongitude: 106.8456,
		DestinationLat:  -6.2088,
		DestinationLong: 106.8456,
		TotalPayment:    500000,
		OrderDate:       time.Now(),
		Shipping:        shiping.Shiping{ /* Sesuaikan dengan data Shiping */ },
	}

	// Create Order
	createdOrder, err := repo.Create(order)
	if err != nil {
		log.Fatalf("Error creating order: %v", err)
	}

	// Output created order
	fmt.Println("Created Order:", createdOrder)

	// Get Order by ID
	fetchedOrder, err := repo.Get(createdOrder.ID)
	if err != nil {
		log.Fatalf("Error fetching order: %v", err)
	}
	fmt.Println("Fetched Order:", fetchedOrder)
}
