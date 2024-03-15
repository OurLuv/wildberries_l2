package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// возвращает строки из файла
func GetLines(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("can't read a file: %s", err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	var lines [][]string
	var words []string
	for scan.Scan() {
		words = strings.Fields(scan.Text())
		lines = append(lines, words)
	}

	return lines
}

// удаляет дубликаты
func DeleteDuplicates(lines [][]string) [][]string {
	mp := make(map[string]struct{})
	var result [][]string

	for _, v := range lines {
		if _, ok := mp[strings.Join(v[:], ",")]; ok {
			continue
		}
		mp[strings.Join(v[:], ",")] = struct{}{}
		result = append(result, v)
	}
	return result
}

// проверяет отсортирован ли слайс
func isSorted(lines [][]string, reverse bool) bool {
	for i := 1; i < len(lines); i++ {
		if reverse {
			if lines[i-1][0] < lines[i][0] {
				return false
			}
		} else {
			if lines[i-1][0] > lines[i][0] {
				return false
			}
		}
	}
	return true
}

// значения с учётом суффиксов
func parseNumericSuffix(s string) (int, string) {
	for i := len(s) - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(s[i])) {
			num, _ := strconv.Atoi(s[:i+1])
			return num, s[i+1:]
		}
	}
	return 0, s
}

func main() {
	// flags
	filepath := flag.String("file", "input.txt", "Path to the file")
	column := *flag.Int("k", 0, "Column to sort by")
	num := *flag.Bool("n", false, "Sort numerically")
	reverse := *flag.Bool("r", false, "Reverse sort")
	unique := *flag.Bool("u", false, "Display unique lines only")
	month := *flag.Bool("m", false, "Sort by month name")
	checkSorted := *flag.Bool("c", false, "Check if data is sorted")
	suffix := *flag.Bool("h", false, "Sort numerically with suffixes")
	flag.Parse()
	lines := GetLines(*filepath)
	if month {
		sort.Slice(lines, func(i, j int) bool {
			date1, err1 := time.Parse("January", lines[i][column])
			date2, err2 := time.Parse("January", lines[j][column])
			if err1 == nil && err2 == nil {
				if reverse {
					return date1.After(date2)
				}
				return date1.Before(date2)
			}
			return false
		})
	} else {
		sort.SliceStable(lines, func(i, j int) bool {
			if num {
				num1, err1 := strconv.Atoi(lines[i][column])
				num2, err2 := strconv.Atoi(lines[j][column])
				if err1 == nil && err2 == nil {
					if suffix {
						num1, suffix1 := parseNumericSuffix(lines[i][column])
						num2, suffix2 := parseNumericSuffix(lines[j][column])
						if num1 != num2 {
							if reverse {
								return num1 > num2
							}
							return num1 < num2
						}
						return suffix1 < suffix2
					}
					if reverse {
						return num1 < num2
					}
					return num1 > num2
				}
				return false
			}
			if reverse {
				return lines[i][column] > lines[j][column]
			}
			return lines[i][column] < lines[j][column]
		})
	}
	if unique {
		lines = DeleteDuplicates(lines)
	}
	if checkSorted {
		if isSorted(lines, reverse) {
			fmt.Println("Data is sorted")
		} else {
			fmt.Println("Data is not sorted")
		}
	}

	fmt.Println(lines)
}
