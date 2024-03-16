package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func download(urlStr string) error {
	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	file, err := os.Create("index.html")
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(body)
	return nil
}

func main() {
	url := flag.String("url", "", "page to download")
	flag.Parse()

	if err := download(*url); err != nil {
		fmt.Printf("error downloading page: %s", err)
	}
}
