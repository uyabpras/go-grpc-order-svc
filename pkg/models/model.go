package models

type Order struct {
	ID        int64 `json:"id" gorm:"primaryKey"`
	Price     int64 `json:"price"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
	UserID    int64 `json:"user_id"`
}
