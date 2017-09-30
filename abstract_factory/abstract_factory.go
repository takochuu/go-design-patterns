package abstract_factory

type factory interface {
	GetSoup() string
	GetMain() string
}

type MizutakiFactory struct {
	factory
}

func NewMizutakiFactory() MizutakiFactory {
	return MizutakiFactory{}
}

func (m MizutakiFactory) GetSoup() string {
	return "Miso"
}

func (m MizutakiFactory) GetMain() string {
	return "Chicken"
}

type HotPot struct {
	Soup string
	Main string
}

func NewHotPot() *HotPot {
	return &HotPot{}
}

func (h *HotPot) AddSoup(soup string) {
	h.Soup = soup
}

func (h *HotPot) AddMain(main string) {
	h.Main = main
}

func (h *HotPot) Print() string {
	return h.Soup + h.Main
}
