package app

type App interface {
	Run()
}

type app struct {
}

func New() App {
	return &app{}
}

func (a app) Run() {

}
