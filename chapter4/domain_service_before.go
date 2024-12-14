package chapter4

// Domain Services are stateless operations within a domain
// that complete a certain activity. Sometimes we cannot find a good way
// to model in an entity or value object; in this cases its a good
// idea to use a domain service.

/// Some reasons to have a domain service
// • The code you are about to write performs a significant piece of
// business logic within one domain
// • You are transforming one domain object into another
// • You are taking the properties of two or more domain objects to calculate a value

type Product struct {
	ID             int
	InStock        bool
	InSomeonesCart bool
}

func (p *Product) CanBeBought() bool {
	return p.InStock && !p.InSomeonesCart
}

type ShoppingCart struct {
	ID          int
	Products    []Product
	IsFull      bool
	MaxCartSize int
}

func (s *ShoppingCart) AddToCart(p Product) bool {
	if s.IsFull {
		return false
	}
	if p.CanBeBought() {
		s.Products = append(s.Products, p)
		return true
	}
	if s.MaxCartSize == len(s.Products) {
		s.IsFull = true
	}
	return true
}
