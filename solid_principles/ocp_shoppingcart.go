package solid_principles

type Item interface {
	Price() float64
}

type ShoppingCart struct {
	Items []Item
}

func (cart *ShoppingCart) AddItem(item Item) {
	cart.Items = append(cart.Items, item)
}

func (cart *ShoppingCart) CalculateTotal() float64 {
	total := 0.0
	for _, item := range cart.Items {
		total += item.Price()
	}
	return total
}
