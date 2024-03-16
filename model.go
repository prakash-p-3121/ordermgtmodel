package ordermgtmodel

import "time"

const (
	OrderCreated              uint = 1
	OrderCancelled            uint = 2
	OrderProcessing           uint = 3
	OrderShipped              uint = 4
	OrderOutForDelivery       uint = 5
	OrderDelivered            uint = 6
	OrderReturnRequested      uint = 7
	OrderReturned             uint = 8
	OrderReplacementRequested uint = 9
	OrderReplaced             uint = 10
)

type Order struct {
	ID           string    `json:"id"`
	IDBitCount   string    `json:"id-bit-count"`
	ProductID    string    `json:"product-id"`
	SellerID     string    `json:"seller-id"`
	BuyerID      string    `json:"buyer-id"`
	ListingID    string    `json:"listing-id"`
	Currency     string    `json:"currency"`
	SellingPrice float64   `json:"selling-price"`
	StateID      uint      `json:"state-id"`
	CreatedAt    time.Time `json:"created-at"`
	UpdatedAt    time.Time `json:"updated-at"`
}
