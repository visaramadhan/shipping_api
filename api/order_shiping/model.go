package order_shiping

import (
	"time"

	"github.com/visaramadhan/shipping_api.git/api/shiping"
)

type OrderShipping struct {
	ID              string          `gorm:"type:varchar(255);primaryKey;unique;not null" json:"id"`
	EcommerceID     string          `gorm:"type:varchar(255);not null" json:"e_commerce_id"`
	ShippingID      string          `gorm:"type:varchar(255);not null" json:"shipping_id" binding:"required"`
	Shipping        shiping.Shiping `gorm:"type:varchar(255); not null" json:"shipping"`
	OriginLongitude float64         `gorm:"type:decimal(10,8);not null" json:"origin_lat" binding:"required"`
	DestinationLat  float64         `gorm:"type:decimal(10,8);not null" json:"destination_lat" binding:"required"`
	DestinationLong float64         `gorm:"type:decimal(11,8);not null" json:"destination_lng" binding:"required"`
	TotalPayment    float64         `gorm:"type:decimal(15,2);not null" json:"total_payment" binding:"required"`
	OrderDate       time.Time       `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"order_date"`
}
