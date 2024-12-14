package chapter3

import (
	"time"

	"github.com/Rhymond/go-money"
)

// Auction is an entity to represent our auction construct
type Auction struct {
	ID            int
	startingPrice money.Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}

// Anemic Entity
type AnemicAuction struct {
	id            int
	startingPrice money.Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}
