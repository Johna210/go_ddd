package chapter3

import (
	"errors"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

// Before trying to cluster our domain models into aggregates, we need to find our
// bounded context's invariants.
// An invariant is a simple rule in our domain that must always be true.
// eg -> we may say in our system for an order to be created we must have the
// item in stock.

// For aggregates, we are looking for transactional consistency,
// not eventual consistency.
// We want any changes to our aggregate to be immediate and atomic.
// So we can think of an aggregate as a transactional consistency boundary.

// Whenever we make changes within our domain, we should ideally only
// modify one aggregate per transaction.

type WalletItem interface {
	GetBalance() (money.Money, error)
}

type Wallet struct {
	id          uuid.UUID // aggregate root and our wallet identity.
	ownerID     uuid.UUID // the ID of the entity that owns the wallet.
	walletItems []WalletItem
}

func (w Wallet) GetWalletBalance() (*money.Money, error) {
	var bal *money.Money
	for _, v := range w.walletItems {
		itemBal, err := v.GetBalance()
		if err != nil {
			return nil, errors.New("failed to get balance")
		}
		bal, err = bal.Add(&itemBal)
		if err != nil {
			return nil, errors.New("failed to increment balance")
		}
	}
	return bal, nil
}
