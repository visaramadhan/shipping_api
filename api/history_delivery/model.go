package history_delivery

import (
	"github.com/visaramadhan/shipping_api.git/api/order_shiping"
)

type History struct {
	ID       string                      `gorm:"type:varchar(255);primaryKey;not null" json:"id_history"`
	OrderID  string                      `gorm:"type:varchar(255);not null" json:"order_id"`
	Order    order_shiping.OrderShipping `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"order"`
	Location string                      `gorm:"type:varchar(255);not null" json:"location"`
	Status   string                      `gorm:"type:varchar(255);not null" json:"status"`
}
