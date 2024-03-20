package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"dev10/server"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type Client struct {
	Conn net.Conn
}

// dial to server
func (c *Client) Dial(address string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return err
	}
	c.Conn = conn
	return nil
}

// put message in connection
func (c *Client) DoEcho(msg string) (string, error) {
	if _, err := c.Conn.Write([]byte(msg)); err != nil {
		return "", err
	}

	buf := make([]byte, 1024)
	_, err := c.Conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			return "", err
		}
	}

	fmt.Println()
	return string(buf), nil
}

func NewClient() *Client {
	return &Client{}
}

func main() {
	go server.StartTCP()
	time.Sleep(time.Second)

	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: go-telnet [--timeout=<duration>] host port")
		os.Exit(1)
	}
	host := args[0]
	port := args[1]

	client := NewClient()
	if err := client.Dial(fmt.Sprintf("%s:%s", host, port), *timeout); err != nil {
		fmt.Printf("error dialing to a server[%s]: %s\n", port, err)
		os.Exit(1)
	}
	defer client.Conn.Close()
	exit := make(chan struct{})
	go func() {
		defer client.Conn.Close()
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Println("Ctrl+D pressed. Closing connection...")
					exit <- struct{}{}
				} else {
					fmt.Println("Error reading from STDIN:", err)
				}
				break
			}
			resp, err := client.DoEcho(text)
			if err != nil {
				fmt.Printf("error [DoMessage]: %s\n", err)
				exit <- struct{}{}
			}
			fmt.Println(resp)
		}
	}()

	<-exit
}
