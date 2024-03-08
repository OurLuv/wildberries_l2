package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// интерфейс для наших стратегий
// стратегии реализовывают похожее поведение
type Strategy interface {
	Route(a, b int)
}

type Navigator struct {
	Strategy
}

// метод установки конкретной стратегии на лету
func (n *Navigator) SetStrategy(s Strategy) {
	n.Strategy = s
}

// первая стратегия - Road
type Road struct{}

func (r *Road) Route(a, b int) {
	// расчёт времени для машины
	t := (b - a) / 60 * 2
	fmt.Printf("Road: %dmins", t)
}

// вторая стратегия - Walk
type Walk struct{}

func (w *Walk) Route(a, b int) {
	// расчёт времени пешком
	t := (b - a - 10) / 5
	fmt.Printf("Walk: %dmins", t)
}

// nav.SetStrategy(walk)
// naw.Route(10, 25)
