package ordermgtmodel

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/restlib"
)

type OrderCreateReq struct {
	ProductID *string  `json:"product-id"`
	SellerID  *string  `json:"seller-id"`
	BuyerID   *string  `json:"buyer-id"`
	ListingID *string  `json:"listing-id"`
	Currency  *string  `json:"currency"`
	Amount    *float64 `json:"selling-price"`
}

func (req *OrderCreateReq) Validate() errorlib.AppError {
	if req.ProductID == nil {
		return errorlib.NewBadReqError("product-id-nil")
	}
	if restlib.TrimAndCheckForEmptyString(req.ProductID) {
		return errorlib.NewBadReqError("product-id-empty")
	}

	if req.SellerID == nil {
		return errorlib.NewBadReqError("seller-id-nil")
	}
	if restlib.TrimAndCheckForEmptyString(req.SellerID) {
		return errorlib.NewBadReqError("seller-id-empty")
	}

	if req.BuyerID == nil {
		return errorlib.NewBadReqError("buyer-id-nil")
	}
	if restlib.TrimAndCheckForEmptyString(req.BuyerID) {
		return errorlib.NewBadReqError("buyer-id-empty")
	}

	if req.ListingID == nil {
		return errorlib.NewBadReqError("listing-id-nil")
	}
	if restlib.TrimAndCheckForEmptyString(req.ListingID) {
		return errorlib.NewBadReqError("listing-id-empty")
	}

	if req.Currency == nil {
		return errorlib.NewBadReqError("currency-nil")
	}
	if restlib.TrimAndCheckForEmptyString(req.Currency) {
		return errorlib.NewBadReqError("currency-empty")
	}

	if req.Amount == nil {
		return errorlib.NewBadReqError("billing-price-nil")
	}

	if *req.Amount <= 0 {
		return errorlib.NewBadReqError("selling-price<=0")
	}

	return nil
}
