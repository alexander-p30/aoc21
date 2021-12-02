package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	increases := 0
	measurements := strings.Split(readReport(), "\n")
	previousWindowMeasurement := math.MinInt32

	for i := len(measurements) - 2; i >= 2; i-- {
		windowMeasurement := 0

		for _, depth := range measurements[(i - 2):(i + 1)] {
			depth, _ := strconv.Atoi(depth)
			windowMeasurement += depth
		}

		if windowMeasurement < previousWindowMeasurement {
			increases++
		}

		previousWindowMeasurement = windowMeasurement
	}

	fmt.Println(increases)
}

func readReport() string {
	pwd, _ := os.Getwd()
	reportFile, _ := os.ReadFile(pwd + "/b/report.txt")

	return string(reportFile)
}
