package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Command interface {
	execute()
}

type Restaurant struct {
	TotalDishes   int
	CleanedDishes int
}

// вынесем наши команды в отдельный структуры - инкапсулируем их
// и реализуем интерфейс Command
type MakeBurgerCommand struct {
	n int
	*Restaurant
}
type CleanDishesCommand struct {
	*Restaurant
}

func (mbc *MakeBurgerCommand) execute() {
	mbc.Restaurant.CleanedDishes -= mbc.n
	fmt.Printf("Made %d burgers\n", mbc.n)
}
func (cdc *CleanDishesCommand) execute() {
	cdc.Restaurant.CleanedDishes = cdc.Restaurant.TotalDishes
	fmt.Println("All dishes are cleaned")
}

// создадим методы Ресторана, чтобы создать экземпляры команд
func (r *Restaurant) MakeBurger(n int) Command {
	return &MakeBurgerCommand{
		n:          n,
		Restaurant: r,
	}
}
func (r *Restaurant) CleanDishes() Command {
	return &CleanDishesCommand{
		Restaurant: r,
	}
}
