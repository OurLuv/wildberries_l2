package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// * cd
func cd(args []string) {
	err := os.Chdir(args[1])
	if err != nil {
		fmt.Println("error changing directory:", err)
	} else {
		fmt.Println("Directory changed")
	}
}

// * echo
func echo(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

// * pwd
func pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current directory:", err)
	} else {
		fmt.Println(dir)
	}
}

// * ps
func ps() {
	cmd := exec.Command("ps")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("error running ps command:", err)
	} else {
		fmt.Println(string(out))
	}
}

// * kill
func kill(args []string) {
	pid := args[1]
	cmd := exec.Command("kill", pid)
	err := cmd.Run()
	if err != nil {
		fmt.Println("error killing process:", err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input:", err)
			continue
		}

		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":
			cd(args)
		case "pwd":
			pwd()
		case "echo":
			echo(args)
		case "kill":
			kill(args)
		case "ps":
			ps()
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
		}
	}
}
