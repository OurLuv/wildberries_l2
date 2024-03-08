package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// создаём интерфейс для наших хендлеров
type HandlerInterface interface {
	SetNext(Handler)
	Handle(int)
}

type Middleware struct {
	Next HandlerInterface
}

// в SetNext мы указываем какой хендлер будет вызван следующим в цепочке
func (m *Middleware) SetNext(handler HandlerInterface) {
	m.Next = handler
}

// метод-обработчик
func (m *Middleware) Handle(data int) {
	fmt.Println("Working with data")
	data += 10
	// вызываем следующий хендлер, передаем в него данные
	m.Next.Handle(data)
}

type Handler struct {
	Next HandlerInterface
}

func (h *Handler) SetNext(handler HandlerInterface) {
	h.Next = handler
}

func (h *Handler) Handle(data int) {
	// выводим полученные данные
	fmt.Printf("Updated data: %d\n", data)
}
