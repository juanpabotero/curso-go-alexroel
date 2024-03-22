package animal

type Animal interface {
	Sonido() string
}

type Dog struct {
	Nombre string
}

func (p *Dog) Sonido() string {
	return "Guau"
}

type Cat struct {
	Nombre string
}

func (g *Cat) Sonido() string {
	return "Miau"
}

func AnimalSound(a Animal) string {
	return a.Sonido()
}
