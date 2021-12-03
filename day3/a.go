package main

import (
	"fmt"
	"os"
	"strings"
)

const diagnosticLineLength = 12
const zero = 48
const one = 49

type PowerConsumtion struct {
	GammaRate   int
	EpsilonRate int
}

func main() {
	diagnosticReport := strings.Split(readReport(), "\n")
	diagnosticSize := len(diagnosticReport) - 1
	diagnosticReport = diagnosticReport[:diagnosticSize]
	powerConsumption := newPowerConsumption()

	var onesCount int
	for i := 0; i < diagnosticLineLength; i++ {
		onesCount = 0
		for j := 0; j < diagnosticSize; j++ {
			if diagnosticReport[j][i] == one {
				onesCount++
			}
		}
		mostFrequentBit, leastFrequentBit := orderBitsByFrequency(onesCount, diagnosticSize)

		updatePowerConsumption(powerConsumption, mostFrequentBit, leastFrequentBit)
	}

	consumption := powerConsumption.GammaRate * powerConsumption.EpsilonRate
	fmt.Printf("𝛾 * ε = %d\n", consumption)
}

func newPowerConsumption() *PowerConsumtion {
	p := new(PowerConsumtion)
	p.GammaRate = 0
	p.EpsilonRate = 0
	return p
}

func updatePowerConsumption(p *PowerConsumtion, mostFrequentBit int, leastFrequentBit int) {
	p.GammaRate = p.GammaRate*2 + mostFrequentBit
	p.EpsilonRate = p.EpsilonRate*2 + leastFrequentBit
}

func orderBitsByFrequency(onesCount int, iteractionCount int) (int, int) {
	if onesCount >= iteractionCount/2 {
		return 1, 0
	}

	return 0, 1
}

func readReport() string {
	pwd, _ := os.Getwd()
	reportFile, _ := os.ReadFile(pwd + "/report_a.txt")

	return string(reportFile)
}
