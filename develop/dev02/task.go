package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unpack(s string) (string, error) {
	var result strings.Builder
	var repeat int
	var esc bool
	for i, r := range s {
		if unicode.IsDigit(r) && !esc {
			repeat, _ = strconv.Atoi(string(r))
			repeat -= 1
			if i == 0 {
				return "", errors.New("некорректная строка")
			}
			if i < len(s)-1 {
				if unicode.IsDigit(rune(s[i+1])) {
					return "", errors.New("некорректная строка")
				}
			}
			result.WriteString(strings.Repeat(string(s[i-1]), repeat))
		} else if r == '\\' && !esc {
			esc = true
		} else {
			result.WriteString(string(r))
			esc = false
		}
	}
	return result.String(), nil
}

func main() {
	str, err := Unpack("d4s5")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
