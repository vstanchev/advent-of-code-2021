package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read everything into a slice.
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// Report any errors from the scanner if it stopped
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	// The oxygen generator and CO2 scrubber rating are found by either the most common or least common bit.
	criterias := map[string]func(length, countOnes int) rune{
		"oxygen": func(countZeros, countOnes int) rune {
			if countOnes >= countZeros {
				return '1'
			}
			return '0'
		},
		"co2": func(countZeros, countOnes int) rune {
			if countOnes >= countZeros {
				return '0'
			}
			return '1'
		},
	}

	ratings := map[string]int64{}

	for rating, criteria := range criterias {
		i := 0
		readings := make([]string, len(input))
		copy(readings, input)
		for len(readings) > 1 {
			fmt.Printf("[%s] Bit position %v", rating, i)
			// Count the 1's at the bit position being considered.
			countOnes := 0
			for _, r := range readings {
				if r[i] == '1' {
					countOnes++
				}
			}
			fmt.Printf(" has %d ones, %d zeros", countOnes, len(readings)-countOnes)

			// Determine the common bit depending on the criteria function.
			cb := criteria(len(readings)-countOnes, countOnes)
			fmt.Printf(" => common bit is %q", cb)

			// Create a new slice with only the readings that have the MCB at the current position.
			var newReadings []string
			for _, r := range readings {
				if rune(r[i]) == cb {
					newReadings = append(newReadings, r)
				}
			}
			fmt.Printf("; Keeping %v \n", newReadings)
			readings = newReadings
			// Move to the next bit position
			i++
		}
		readings[0] = fmt.Sprintf("0b%s", readings[0])
		fmt.Printf("%v\n", readings)
		ratings[rating], _ = strconv.ParseInt(readings[0], 0, 0)
	}

	fmt.Printf("Ratings are %v\n", ratings)
	fmt.Printf("Result is %d\n", ratings["oxygen"]*ratings["co2"])
}
