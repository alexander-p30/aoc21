package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	HorizontalPosition int
	Depth              int
}

func main() {
	sub := new(Submarine)
	commands := strings.Split(readReport(), "\n")
	commands = commands[:(len(commands) - 1)]

	for _, command := range commands {
		resolveCommand(sub, strings.Split(command, " "))
	}

	fmt.Println(sub.Depth * sub.HorizontalPosition)
}

func readReport() string {
	pwd, _ := os.Getwd()
	reportFile, _ := os.ReadFile(pwd + "/a/report.txt")

	return string(reportFile)
}

func resolveCommand(sub *Submarine, command []string) {
	amount, _ := strconv.Atoi(command[1])

	switch command[0] {
	case "forward":
		sub.HorizontalPosition = sub.HorizontalPosition + amount
	case "down":
		sub.Depth = sub.Depth + amount
	case "up":
		sub.Depth = sub.Depth - amount
	}
}
