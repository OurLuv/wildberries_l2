package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func GetContent(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	buf := strings.Builder{}
	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		buf.Write(data[:n])
	}
	return buf.String(), nil
}

// возвращает строки из файла
func GetLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	var lines []string
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}

	return lines, nil
}

func main() {
	filepath := *flag.String("file", "data.txt", "Path to the input file")
	pattern := *flag.String("pattern", "", "Pattern to search for")
	after := *flag.Int("A", 0, "Print +N lines after match")
	before := *flag.Int("B", 0, "Print +N lines before match")
	context := *flag.Int("C", 0, "Print ±N lines around match")
	count := *flag.Bool("c", false, "Print count of matching lines")
	ignore := *flag.Bool("i", false, "Ignore case")
	invert := *flag.Bool("v", false, "Invert match")
	fixed := *flag.Bool("F", false, "Exact match with string, not pattern")
	num := *flag.Bool("n", false, "Print line numbers")
	flag.Parse()

	lines, err := GetLines(filepath)
	if err != nil {
		log.Fatal(err)
	}
	var countLines []string

	for i, v := range lines {
		match := false
		if fixed {
			match = strings.Contains(v, pattern)
		} else {
			re := regexp.MustCompile(pattern)
			match = re.MatchString(v)
		}
		if (ignore && match) || (!ignore && strings.Contains(strings.ToLower(v), strings.ToLower(pattern))) {
			if invert {
				continue
			}
			if count {
				continue
			}
			if before > 0 {
				for j := i - before; j < i; j++ {
					if j >= 0 {
						fmt.Println(formatLine(j, lines[j], num))
					}
				}
			}
			fmt.Println(formatLine(i, v, num))
			if after > 0 {
				for j := i + 1; j <= i+after; j++ {
					if j < len(lines) {
						fmt.Println(formatLine(j, lines[j], num))
					}
				}
			}
			if context > 0 {
				start := i - context
				if start < 0 {
					start = 0
				}
				end := i + context
				if end >= len(lines) {
					end = len(lines) - 1
				}
				for j := start; j <= end; j++ {
					fmt.Println(formatLine(j, lines[j], num))
				}
			}
			countLines = append(countLines, v)
		} else if invert {
			fmt.Println(formatLine(i, v, num))
		}
	}

	if count {
		fmt.Println("Count of matching lines:", len(countLines))
	}
}

func formatLine(num int, line string, print bool) string {
	if print {
		return fmt.Sprintf("%d:%s", num+1, line)
	}
	return line
}
