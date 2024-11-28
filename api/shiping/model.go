package shiping

type Shiping struct {
	ID          string  `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	Name        string  `gorm:"type:varchar(255);not null" json:"name" binding:"required,alphanum"`
	Price       float64 `gorm:"type:decimal(10,2);not null" json:"price" binding:"required,gt=0"`
	Address     string  `gorm:"type:varchar(255);not null" json:"address" binding:"required"`
	Destination string  `gorm:"type:varchar(255);not null" json:"destination" binding:"required"`
}
