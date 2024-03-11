package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// интерфейс состояния с методом для публикации статьи
type State interface {
	publish()
}

// состояние статьи - черновик
type StateDraft struct{}

func (sd *StateDraft) publish() {
	fmt.Println("Published to moderator")
}

// состояние статьи - на модерации
type StateMod struct{}

func (sm *StateMod) publish() {
	fmt.Println("Published out")
}

type Article struct {
	State
}

// метод для установки конкретного состояния статьи
func (a *Article) SetState(s State) {
	a.State = s
}
