package decorator

type icecream interface {
	GetName() string
	GetCost() int
}

type Matcha struct {
	icecream
}

func NewMatcha() *Matcha {
	return &Matcha{}
}

func (m *Matcha) GetName() string {
	return "matcha"
}

func (m *Matcha) GetCost() int {
	return 300
}

type CashewNuts struct {
	icecream
	parent icecream
}

func NewCashewNuts(ice icecream) *CashewNuts {
	return &CashewNuts{
		parent: ice,
	}
}

func (m *CashewNuts) GetName() string {
	return m.parent.GetName() + "cashewnuts"
}

func (m *CashewNuts) GetCost() int {
	return m.parent.GetCost() + 100
}
