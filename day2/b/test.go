package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Directive string
	Count     int
}

type Submarine struct {
	HorizontalPosition int
	Depth              int
	Aim                int
}

func main() {
	sub := newSubmarine()
	commands := strings.Split(readReport(), "\n")
	commands = commands[:(len(commands) - 1)]

	for _, command := range commands {
		resolveCommand(sub, newCommand(command))
	}

	fmt.Println(sub.Depth * sub.HorizontalPosition)
}

func readReport() string {
	pwd, _ := os.Getwd()
	reportFile, _ := os.ReadFile(pwd + "/b/report.txt")

	return string(reportFile)
}

func newSubmarine() *Submarine {
	sub := new(Submarine)
	sub.HorizontalPosition = 0
	sub.Depth = 0
	sub.Aim = 0

	return sub
}

func newCommand(commandString string) *Command {
	commandValues := strings.Split(commandString, " ")
	cmd := new(Command)
	cmd.Directive = commandValues[0]
	cmd.Count, _ = strconv.Atoi(commandValues[1])

	return cmd
}

func resolveCommand(sub *Submarine, cmd *Command) {
	switch cmd.Directive {
	case "forward":
		sub.HorizontalPosition = sub.HorizontalPosition + cmd.Count
		sub.Depth = sub.Depth + (cmd.Count * sub.Aim)
	case "down":
		sub.Aim = sub.Aim + cmd.Count
	case "up":
		sub.Aim = sub.Aim - cmd.Count
	}
}
