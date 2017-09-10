package composite

type AmountInterface interface {
	GetAmount() int
}

type list struct {
	orders []AmountInterface
}

func Newlist() *list {
	return &list{}
}

func (l *list) Add(i AmountInterface) {
	l.orders = append(l.orders, i)
}

func (l *list) GetAmount() int {
	amount := 0
	for _, v := range l.orders {
		amount += v.GetAmount()
	}
	return amount
}

type productA struct {
	amount int
	AmountInterface
}

func NewProductA() *productA {
	return &productA{
		amount: 100,
	}
}

func (p *productA) GetAmount() int {
	return p.amount
}

type productB struct {
	amount int
	AmountInterface
}

func NewProductB() *productB {
	return &productB{
		amount: 200,
	}
}

func (p *productB) GetAmount() int {
	return p.amount
}
