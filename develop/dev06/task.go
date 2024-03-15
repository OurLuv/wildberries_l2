package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fields := flag.String("f", "", "Fields")
	delimiter := flag.String("d", "\t", "Delimiter")
	separated := flag.Bool("s", false, "Separated")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		columns := strings.Split(line, *delimiter)

		if *fields == "" {
			fmt.Println(line)
		} else {
			var selectedFields []string
			fieldNums := strings.Split(*fields, ",")
			for _, v := range fieldNums {
				index, _ := strconv.Atoi(v)
				index -= 1
				if index >= 0 && index < len(columns) {
					selectedFields = append(selectedFields, columns[index])
				}
			}
			fmt.Println(strings.Join(selectedFields, *delimiter))
		}
	}
}
