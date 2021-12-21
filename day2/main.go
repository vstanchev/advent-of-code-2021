package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	all, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	partOne(string(all))
	partTwo(string(all))
}

func partOne(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	sums := map[string]int{
		"forward": 0,
		"down":    0,
		"up":      0,
	}

	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")
		v, _ := strconv.Atoi(cmd[1])
		sums[cmd[0]] += v
	}

	// Report any errors from the scanner if it stopped
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	// Calculate the result by removing the "up" sum from the "down" sum and multiplying by "forward"
	fmt.Printf("Part 1 result %v\n", sums["forward"]*(sums["down"]-sums["up"]))
}

func partTwo(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	aim := 0
	depth := 0
	horizontal := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		cmd := strings.Split(scanner.Text(), " ")
		v, _ := strconv.Atoi(cmd[1])
		switch cmd[0] {
		case "down":
			aim += v
		case "up":
			aim -= v
		case "forward":
			depth += aim * v
			horizontal += v
		}
	}

	// Report any errors from the scanner if it stopped
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	// Calculate the result by removing the "up" sum from the "down" sum and multiplying by "forward"
	fmt.Printf("Part 2 result depth: %v, horizontal: %v, multiplied: %v\n", depth, horizontal, depth*horizontal)
}
