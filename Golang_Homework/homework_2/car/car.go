package car

type car struct {
	name  string
	price int
	color string
}

//New struct new function
func New(Name string, Price int, Color string) *car {
	return &car{
		name:  Name,
		price: Price,
		color: Color,
	}
}

func (c *car) GetName() string {
	return c.name
}

func (c *car) Add(prices ...int) int {
	for _, itemPrice := range prices {
		c.price += itemPrice
	}
	return c.GetPrice()
}

func (c *car) Remove(prices ...int) int {
	for _, itemPrice := range prices {
		c.price -= itemPrice
	}
	return c.GetPrice()
}

func (c *car) GetPrice() int {
	return c.price
}

func (c *car) GetColor() string {
	return c.color
}

func (c *car) UpdateColor(col string) (result string) {
	c.color = col
	result = "Sucess! Car's color: " + c.GetColor()
	return result
}
