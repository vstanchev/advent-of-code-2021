package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	measurements := getMeasurements(scanner)
	diffCount := 0

	// Init first window sum.
	prevWindow := measurements[:3]
	prevSum := 0
	for _, v := range prevWindow {
		prevSum += v
	}
	// Remove first value as it won't be needed for the next window.
	measurements = measurements[1:]

	// Loop through all measurements
	for window := measurements[:3]; len(measurements) >= 3; window = measurements[:3] {
		sum := 0
		for _, v := range window {
			sum += v
		}
		if sum > prevSum {
			diffCount++
		}

		// Prepare for next iteration by removing one value and remembering the sum of the window.
		measurements = measurements[1:]
		prevSum = sum
	}

	fmt.Printf("The number of times depth increased is %d", diffCount)

}

func getMeasurements(scanner *bufio.Scanner) []int {
	var measurements []int
	for scanner.Scan() {
		text := scanner.Text()
		depth, err := strconv.Atoi(text)
		if err != nil {
			log.Printf("%q is not a number", text)
			continue
		}
		measurements = append(measurements, depth)
	}

	// Report any errors from the scanner if it stopped
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return measurements
}
