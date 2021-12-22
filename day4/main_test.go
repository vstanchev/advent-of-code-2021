package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func Test_newBoard(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want Board
	}{
		{
			"1x2 board",
			args{input: "22\n 5"},
			Board{
				[][]int{
					{22},
					{5},
				},
				[][]bool{
					{false},
					{false},
				},
				0,
			},
		},
		{
			"5x5 board",
			args{input: "22 13 17 11  0\n 8  2 23  4 24\n21  9 14 16  7\n 6 10  3 18  5\n 1 12 20 15 19"},
			Board{
				[][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				[][]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
				},
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBoard(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBoards(t *testing.T) {
	type args struct {
		scanner *bufio.Scanner
	}
	tests := []struct {
		name string
		args args
		want []Board
	}{
		{
			"2 1x2 boards",
			args{scanner: bufio.NewScanner(strings.NewReader("22\n 5\n\n11\n 4\n"))},
			[]Board{
				{
					[][]int{
						{22},
						{5},
					},
					[][]bool{
						{false},
						{false},
					},
					0,
				},
				{
					[][]int{
						{11},
						{4},
					},
					[][]bool{
						{false},
						{false},
					},
					0,
				},
			},
		},
		{
			"3 5x5 boards",
			args{scanner: bufio.NewScanner(strings.NewReader("22 13 17 11  0\n 8  2 23  4 24\n21  9 14 16  7\n 6 10  3 18  5\n 1 12 20 15 19\n\n 3 15  0  2 22\n 9 18 13 17  5\n19  8  7 25 23\n20 11 10 24  4\n14 21 16 12  6\n\n14 21 17 24  4\n10 16 15  9 19\n18  8 23 26 20\n22 11 13  6  5\n 2  0 12  3  7"))},
			[]Board{
				{
					[][]int{
						{22, 13, 17, 11, 0},
						{8, 2, 23, 4, 24},
						{21, 9, 14, 16, 7},
						{6, 10, 3, 18, 5},
						{1, 12, 20, 15, 19},
					},
					[][]bool{
						{false, false, false, false, false},
						{false, false, false, false, false},
						{false, false, false, false, false},
						{false, false, false, false, false},
						{false, false, false, false, false},
					},
					0,
				},
				{
					[][]int{
						{3, 15, 0, 2, 22},
						{9, 18, 13, 17, 5},
						{19, 8, 7, 25, 23},
						{20, 11, 10, 24, 4},
						{14, 21, 16, 12, 6},
					},
					[][]bool{
						{false, false, false, false, false},
						{false, false, false, false, false},
						{false, false, false, false, false},
						{false, false, false, false, false},
						{false, false, false, false, false},
					},
					0,
				},
				{
					[][]int{
						{14, 21, 17, 24, 4},
						{10, 16, 15, 9, 19},
						{18, 8, 23, 26, 20},
						{22, 11, 13, 6, 5},
						{2, 0, 12, 3, 7},
					},
					[][]bool{
						{false, false, false, false, false},
						{false, false, false, false, false},
						{false, false, false, false, false},
						{false, false, false, false, false},
						{false, false, false, false, false},
					},
					0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBoards(tt.args.scanner); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBoards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_MarkNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		b    Board
		args args
		want Board
	}{
		{
			"1x2 board mark",
			Board{
				[][]int{
					{22},
					{5},
				},
				[][]bool{
					{false},
					{false},
				},
				0,
			},
			args{num: 5},
			Board{
				[][]int{
					{22},
					{5},
				},
				[][]bool{
					{false},
					{true},
				},
				5,
			},
		},
		{
			"1x2 board no mark",
			Board{
				[][]int{
					{22},
					{5},
				},
				[][]bool{
					{false},
					{false},
				},
				0,
			},
			args{num: 9},
			Board{
				[][]int{
					{22},
					{5},
				},
				[][]bool{
					{false},
					{false},
				},
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.b.marks == nil {
				tt.b.ResetMarks()
			}
			tt.b.MarkNumber(tt.args.num)
			if !reflect.DeepEqual(tt.b, tt.want) {
				t.Errorf("MarkNumber() = %v, want %v", tt.b, tt.want)
			}
		})
	}
}

func TestBoard_IsWinning(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want bool
	}{
		{
			"3x3 winning row",
			Board{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				[][]bool{
					{true, true, true},
					{false, false, false},
					{false, false, false},
				},
				0,
			},
			true,
		},
		{
			"3x3 winning first column",
			Board{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				[][]bool{
					{true, false, false},
					{true, false, false},
					{true, false, false},
				},
				0,
			},
			true,
		},
		{
			"3x3 winning last column",
			Board{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				[][]bool{
					{false, false, true},
					{false, false, true},
					{false, false, true},
				},
				0,
			},
			true,
		},
		{
			"3x3 not winning diagonal",
			Board{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				[][]bool{
					{false, false, true},
					{true, true, false},
					{true, false, true},
				},
				0,
			},
			false,
		},
		{
			"3x3 not winning",
			Board{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				[][]bool{
					{true, false, false},
					{false, true, true},
					{true, false, false},
				},
				0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.IsWinning(); got != tt.want {
				t.Errorf("IsWinning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Score(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		want int
	}{
		{
			"3x3 winning row",
			Board{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				[][]bool{
					{true, true, true},
					{false, false, false},
					{false, false, false},
				},
				2,
			},
			(4 + 5 + 6 + 7 + 8 + 9) * 2,
		},
		{
			"3x3 winning first column",
			Board{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				[][]bool{
					{true, false, false},
					{true, false, false},
					{true, false, false},
				},
				4,
			},
			(2 + 3 + 5 + 6 + 8 + 9) * 4,
		},
		{
			"3x3 winning last column",
			Board{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				[][]bool{
					{false, false, true},
					{false, false, true},
					{false, false, true},
				},
				6,
			},
			(1 + 2 + 4 + 5 + 7 + 8) * 6,
		},
		{
			"3x3 not winning diagonal",
			Board{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				[][]bool{
					{false, false, true},
					{true, true, false},
					{true, false, true},
				},
				0,
			},
			0,
		},
		{
			"3x3 not winning",
			Board{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				[][]bool{
					{true, false, false},
					{false, true, true},
					{true, false, false},
				},
				0,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Score(); got != tt.want {
				t.Errorf("Score() = %v, want %v", got, tt.want)
			}
		})
	}
}
