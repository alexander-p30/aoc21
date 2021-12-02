package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	report, _ := os.ReadFile(pwd + "/report.txt")
	previousDepth := math.MaxInt32
	increases := 0

	for _, depth := range strings.Split(string(report), "\n") {
		depth, _ := strconv.Atoi(depth)

		if depth > previousDepth {
			increases++
		}

		previousDepth = depth
	}

	fmt.Println(increases)
}
