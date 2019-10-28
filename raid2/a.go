package main

import (
	"fmt"
	"os"
)

func print(table [10][10]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(table[i][j])
			fmt.Print(" ")
		}
		fmt.Print(" \n")
	}
	fmt.Print(" \n")
}
func checkcol(num int, col int, table [10][10]int) bool {
	for i := 0; i < 9; i++ {
		if table[i][col] == num {
			return true
		}
	}
	return false

}
func checkrow(num int, row int, table [10][10]int) bool {
	for _, xnum := range table[row] {
		if xnum == num {
			return true
		}
	}
	return false
}

func checksquare(num int, row int, col int, table [10][10]int) bool {
	c := col / 3
	r := row / 3
	c = c * 3
	r = r * 3
	for i := r; i < r+3; i++ {
		for j := c; j < c+3; j++ {
			if table[i][j] == num {
				return true
			}
		}
	}
	return false
}

func checknumber(num int, row int, col int, table [10][10]int) bool {
	if checkrow(num, row, table) == true {
		return true
	} else if checksquare(num, row, col, table) == true {
		return true
	} else if checkcol(num, col, table) == true {
		return true
	} else {
		return false
	}
}
func checkempty(table [10][10]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
func findempty(table [10][10]int) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}
func solvesudoku(table [10][10]int) bool {

	var row, col = findempty(table)
	if row == -1 && col == -1 {
		print(table)
		return true
	} else {
		for n := 0; n <= 9; n++ {
			if checknumber(n, row, col, table) == false {
				table[row][col] = n
				if solvesudoku(table) == true {
					return true
				}
				table[row][col] = 0
			}
		}
		return false

	}
}

func main() {
	count := 0
	for range os.Args {
		count++
	}
	fmt.Println(count)
	var sudoku [10][10]int
	if count == 10 {
		for i := 1; i < count; i++ {
			check := []rune(os.Args[i])
			ln := 0
			for j, ch := range check {

				if ln == 10 {
					fmt.Println("Error")
					return
				}
				if ch >= '0' && ch <= '9' {
					sudoku[i-1][j] = int(ch) - 48
				} else if ch == '.' {
					sudoku[i-1][j] = 0
				} else {

					fmt.Println("Error")
					return
				}
				ln++
			}
		}
		print(sudoku)
		solvesudoku(sudoku)

	} else {
		fmt.Println("Error")
	}
}
