package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func findAnagram(words []string) map[string][]string {
	result := make(map[string][]string)
	lookup := make(map[string]string)

	for _, v := range words {
		word := strings.ToLower(v)
		mask := getChars(word)
		if _, ok := lookup[mask]; !ok {
			lookup[mask] = word
		} else {
			result[lookup[mask]] = append(result[lookup[mask]], word)
		}
	}
	return result
}

func getChars(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s[:], "")
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "машина"}
	result := findAnagram(words)
	for k, v := range result {
		sort.Strings(v)
		fmt.Printf("[%s]: { ", k)
		for _, w := range v {
			fmt.Printf("%s, ", w)
		}
		fmt.Print("}\n")
	}
}
