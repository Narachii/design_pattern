package main

type mocha struct {
	beverage beverage
}

func (m *mocha) description() string {
	return m.beverage.description() + ", Mocha"
}

func (m *mocha) cost() float64 {
	return 0.20 + m.beverage.cost()
}

type milk struct {
	beverage beverage
}

func (m *milk) description() string {
	return m.beverage.description() + ", Milk"
}

func (m *milk) cost() float64 {
	return 0.1 + m.beverage.cost()
}

type soy struct {
	beverage beverage
}

func (s *soy) description() string {
	return s.beverage.description() + ", Soy"
}

func (s *soy) cost() float64 {
	return 0.15 + s.beverage.cost()
}

type whip struct {
	beverage beverage
}

func (w *whip) description() string {
	return w.beverage.description() + ", Whip"
}

func (w *whip) cost() float64 {
	return 0.1 + w.beverage.cost()
}
