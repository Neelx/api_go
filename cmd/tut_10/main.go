package main //generics

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMaker string
	carModel string
	engine   T
}

func main() {
	_ = car[gasEngine]{
		carMaker: "Toyota",
		carModel: "Corolla",
		engine: gasEngine{
			gallons: 10,
			mpg:     30,
		},
	}
	_ = car[electricEngine]{
		carMaker: "Tesla",
		carModel: "Model S",
		engine: electricEngine{
			kwh:   100,
			mpkwh: 30,
		},
	}
}
