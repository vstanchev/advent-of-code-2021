package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Counting "1" bits and the total number of records will help us count the most/least common bits.
	count := 0
	countOnes := make([]int, 12)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			if line[i] == '1' {
				countOnes[i]++
			}
		}
		count++
	}
	fmt.Printf("%v %d\n", countOnes, count)

	// Build the gama and epsilon rate binary string out of the most/least common bits.
	gamaRateStr := "0b"
	epsilonRateStr := "0b"
	for _, bitCount := range countOnes {
		if bitCount > count/2 {
			gamaRateStr += "1"
			epsilonRateStr += "0"
		} else {
			gamaRateStr += "0"
			epsilonRateStr += "1"
		}
	}
	// Parse gama rate and epsilon rate into integers
	gamaRate, err := strconv.ParseInt(gamaRateStr, 0, 0)
	if err != nil {
		panic(err)
	}

	epsilonRate, err := strconv.ParseInt(epsilonRateStr, 0, 0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Gama    rate is %d (%b)\n", gamaRate, gamaRate)
	fmt.Printf("Epsilon rate is %d (%012b)\n", epsilonRate, epsilonRate)
	fmt.Printf("Power consumption is %d", gamaRate*epsilonRate)

	// Report any errors from the scanner if it stopped
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}
