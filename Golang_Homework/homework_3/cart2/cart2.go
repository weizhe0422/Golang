package cart2

type cart2 struct {
	items      []string
	prices     []int
	totalPrice float32
}

//New Initialize the cart obkect
func New(item string, price int) *cart2 {
	return &cart2{
		items:      []string{item},
		prices:     []int{price},
		totalPrice: 0.0,
	}
}

func (c *cart2) Process(moneys ...int) {
	for _, money := range moneys {
		c.prices = append(c.prices, money)
	}

	totalMoney := 0
	for _, countMoney := range c.prices {
		totalMoney += countMoney
	}

	if totalMoney >= 1000 {
		totalMoney -= 200
	}

	c.totalPrice = float32(totalMoney) * 0.7
}

func (c *cart2) GetTotal() float32 {
	return c.totalPrice
}
