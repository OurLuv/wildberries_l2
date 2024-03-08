package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Объект, который мы хотим построить
type Car struct {
	color         string
	engineType    string
	hasSunroof    bool
	hasNavigation bool
}

// CarBuilder это интерфейс для постройки частей Car
type CarBuilder interface {
	SetColor(color string) CarBuilder
	SetEngineType(engineType string) CarBuilder
	SetSunroof(hasSunroof bool) CarBuilder
	SetNavigation(hasNavigation bool) CarBuilder
	Build() *Car
}

// Создаём новый билдер
func NewCarBuilder() CarBuilder {
	return &carBuilder{
		car: &Car{}, // инициализируем атрибут Car
	}
}

// carBuilder структура билдера, который будет реализовывать интерфейс
type carBuilder struct {
	car *Car
}

func (cb *carBuilder) SetColor(color string) CarBuilder {
	cb.car.color = color
	return cb
}

func (cb *carBuilder) SetEngineType(engineType string) CarBuilder {
	cb.car.engineType = engineType
	return cb
}

func (cb *carBuilder) SetSunroof(hasSunroof bool) CarBuilder {
	cb.car.hasSunroof = hasSunroof
	return cb
}

func (cb *carBuilder) SetNavigation(hasNavigation bool) CarBuilder {
	cb.car.hasNavigation = hasNavigation
	return cb
}

func (cb *carBuilder) Build() *Car {
	return cb.car
}

type Director struct {
	builder CarBuilder
}

func (d *Director) ConstructCar(color, engineType string, hasSunroof, hasNavigation bool) *Car {
	// пример использования builder
	d.builder.SetColor(color).
		SetEngineType(engineType).
		SetSunroof(hasSunroof).
		SetNavigation(hasNavigation)

	return d.builder.Build()
}
