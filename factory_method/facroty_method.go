package factory_method

type factory interface {
	GetName() string
}

type Account struct {
	factory
	Name string
}

func (a Account) GetName() string {
	return a.Name
}

func NewAccount(n string) factory {
	return &Account{
		Name: n,
	}
}
