package main 

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64{
	return 2 * math.Pi * c.raio
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

type figura interface {
	area() float64
}

func main() {
	circulo1 := circulo{
		raio: 32.2,
	}
	quadrado1 := quadrado{
		lado: 5,
	}

	fmt.Println(info(quadrado1), info(circulo1))
}

func info(f figura) float64 {
	return f.area()
}