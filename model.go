package ordermgtmodel

import "time"

const (
	OrderReceived             uint = 1
	OrderCancelled            uint = 2
	OrderShipped              uint = 3
	OrderOutForDelivery       uint = 4
	OrderDelivered            uint = 5
	OrderReturnRequested      uint = 6
	OrderReturned             uint = 7
	OrderReplacementRequested uint = 8
	OrderReplaced             uint = 9
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
