package cart1

type cart1 struct {
	items      []string
	prices     []int
	totalPrice float32
}

//New Initialize the cart obkect
func New(item string, price int) *cart1 {
	return &cart1{
		items:      []string{item},
		prices:     []int{price},
		totalPrice: 0.0,
	}
}

func (c *cart1) Process(moneys ...int) {
	for _, money := range moneys {
		c.prices = append(c.prices, money)
	}

	totalMoney := 0
	for _, countMoney := range c.prices {
		totalMoney += countMoney
	}

	if totalMoney >= 500 {
		totalMoney -= 50
	}

	c.totalPrice = float32(totalMoney) * 0.8
}

func (c *cart1) GetTotal() float32 {
	return c.totalPrice
}
