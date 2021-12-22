package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Board represents a 5x5 bingo board that has marked/unmarked numbers
// Example:
//     22 13 17 11  0
//      8  2 23  4 24
//     21  9 14 16  7
//      6 10  3 18  5
//      1 12 20 15 19
type Board struct {
	rows    [][]int
	marks   [][]bool
	lastNum int
}

func main() {
	part1()
	//part2()
}

func part1() {
	fileContents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(fileContents))

	// First row is the order of draw numbers.
	scanner.Scan()
	drawStr := strings.Split(scanner.Text(), ",")
	var draw []int
	for _, s := range drawStr {
		num, _ := strconv.Atoi(s)
		draw = append(draw, num)
	}
	fmt.Printf("Drawn numbers are %v\n\n", draw)

	// Get rid of empty line
	scanner.Scan()

	// Play bingo with drawn numbers and boards.
	boards := getBoards(scanner)
	fmt.Printf("Playing with %d boards\n", len(boards))
	var winningBoard Board
	bingo := false
	for _, num := range draw {
		for _, board := range boards {
			if board.MarkNumber(num) && board.IsWinning() {
				winningBoard = board
				bingo = true
				break
			}
		}
		if bingo {
			break
		}
	}

	fmt.Println("Winning board is")
	winningBoard.PrintBoard()

	fmt.Printf("Score is %d!\n", winningBoard.Score())

	// Report any errors from the scanner if it stopped
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

// Read and parse all boards.
func getBoards(scanner *bufio.Scanner) []Board {
	var boards []Board
	var boardLines strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		// There is an empty line between boards
		if len(line) == 0 {
			boards = append(boards, newBoard(boardLines.String()))
			boardLines.Reset()
			continue
		}
		boardLines.WriteString(line)
		// Add the new line because scanner.Text() does not return it.
		boardLines.WriteRune('\n')
	}
	boards = append(boards, newBoard(boardLines.String()))
	return boards
}

func newBoard(input string) Board {
	var board Board
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		all := strings.TrimSpace(strings.ReplaceAll(line, "  ", " "))
		nums := strings.Split(all, " ")
		var boardRow []int
		for _, numStr := range nums {
			num, _ := strconv.Atoi(numStr)
			boardRow = append(boardRow, num)
		}
		board.rows = append(board.rows, boardRow)
	}
	board.ResetMarks()
	return board
}

func (b *Board) ResetMarks() {
	b.marks = make([][]bool, len(b.rows))
	for i, row := range b.rows {
		b.marks[i] = make([]bool, len(row))
		for j := range row {
			b.marks[i][j] = false
		}
	}
}

// MarkNumber marks a number on a board as drawn.
func (b *Board) MarkNumber(num int) bool {
	for i, row := range b.rows {
		for j, col := range row {
			if col == num {
				b.marks[i][j] = true
				b.lastNum = num
				return true
			}
		}
	}
	return false
}

// IsWinning returns true if the board has a row or a column of marked numbers.
func (b *Board) IsWinning() bool {
	colMarks := make(map[int]bool)
	// Iterate over every row to check if the row is winning and record any marked columns.
	for _, row := range b.marks {
		rowWins := true
		for j, col := range row {
			// Both row and column are not winning in case of an unmarked number.
			if !col {
				rowWins = false
				colMarks[j] = false
				continue
			}

			// Mark column as winning only if it hasn't been marked until now.
			if _, ok := colMarks[j]; !ok {
				colMarks[j] = true
			}
		}

		if rowWins {
			return true
		}
	}

	// Board wins because a column has been marked on every row.
	for _, wins := range colMarks {
		if wins {
			return true
		}
	}

	return false
}

// PrintBoard prints the board on standard output.
func (b Board) PrintBoard() {
	for i, row := range b.rows {
		for j, col := range row {
			if b.marks[i][j] {
				fmt.Printf(" %02d*", col)
			} else {
				fmt.Printf(" %02d ", col)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Score calculates and returns the board score by summing all unmarked numbers
// and multiplying the sum by the last called number.
func (b *Board) Score() int {
	sum := 0
	for i, row := range b.marks {
		for j, col := range row {
			if !col {
				sum += b.rows[i][j]
			}
		}
	}

	return sum * b.lastNum
}
