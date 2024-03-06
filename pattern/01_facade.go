package D

import (
	"errors"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Далее описана сложная структура получения баланса пользователя
// User -> Card -> Bank -> Bank.CardBalance
type Bank struct {
	Name        string
	CardBalance map[string]float64
}

func (b *Bank) GetBalance(cardNumber string) float64 {
	balance := b.CardBalance[cardNumber]
	return balance
}

type Card struct {
	Name string
	Bank *Bank
}

func (c *Card) GetBalance() float64 {
	balance := c.Bank.GetBalance(c.Name)
	return balance
}

type User struct {
	Name         string
	IsSubscribed bool
	Card         *Card
}

func (u *User) GetBalance() float64 {
	balance := u.Card.GetBalance()
	return balance
}

// Структура Spotify будет выступать фасадом
type Spotify struct {
	Subscription float64
}

// В программе нам будет достаточно вызвать метод Subscribe,
// который хранит под копотом всю сложную систему, описанную выше
func (s *Spotify) Subscribe(user User) error {
	if user.GetBalance() < s.Subscription {
		return errors.New("Недостаточно средств")
	}
	user.IsSubscribed = true
	return nil
}
