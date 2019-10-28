package main

import (
	"fmt"
	"os"
)

var a [9][9]int
var ban [10][9][9]int
var solution, err int
var can bool

func BanRow(x, nb int) {
	for j := 0; j < 9; j++ {
		ban[nb][x][j] = 1
	}
}
func BanColumn(y, nb int) {
	for i := 0; i < 9; i++ {
		ban[nb][i][y] = 1
	}
}
func BanSquare(x, y, nb int) {
	for i := x / 3 * 3; i < (x/3+1)*3; i++ {
		for j := y / 3 * 3; j < (y/3+1)*3; j++ {
			ban[nb][i][j] = 1
		}
	}
}
func Conv(x rune) int {
	cur := 1
	for i := '1'; i < x; i++ {
		cur++
	}
	return cur
}
func Check(cur []string) bool {
	if len(cur) != 9 {
		return false
	}
	for x, c := range cur {
		if len(c) != 9 {
			return false
		}
		for y, j := range c {
			if j >= '1' && j <= '9' {
				a[x][y] = Conv(j)
			} else if j != '.' {
				return false
			}
		}
	}
	return true
}
func Initialization() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if a[i][j] > 0 {
				if ban[a[i][j]][i][j] == 1 {
					err = 1
					return
				}
				BanRow(i, a[i][j])
				BanColumn(j, a[i][j])
				BanSquare(i, j, a[i][j])
			}
		}
	}
}
func Show() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(a[i][j])
			if j != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func BruteForce(step int) {
	if ans > 1 {
		return
	}
	if step == -1 {
		ans++
		if can {
			Show()
		}
		return
	}
	var x, y int
	x = step / 9
	y = step % 9
	if a[x][y] == 0 {
		for j := 1; j <= 9; j++ {
			if ban[j][x][y] == 0 {
				q := ban
				a[x][y] = j
				BanRow(x, j)
				BanColumn(y, j)
				BanSquare(x, y, j)
				BruteForce(step - 1)
				a[x][y] = 0
				ban = q
			}
		}
	} else {
		BruteForce(step - 1)
	}
}
func main() {
	arg := os.Args[1:]
	if Check(arg) {
		Initialization()
		if err == 0 {
			BruteForce(80)
		}
		if ans == 1 {
			can = true
			BruteForce(80)
		} else {
			fmt.Print("Error\n")
		}

	} else {
		fmt.Print("Error\n")
	}
}

/*
Coding 23:11:19
./Raid2 "......68." "....73..9" "3.9....45" "49......." "8.3.5.9.2" ".......36" "96....3.8" "7..68...." ".28......"
9 7 8 1 6 3 5 4 2
3 2 1 5 9 4 7 8 6
4 6 5 7 8 2 3 1 9
7 5 2 6 4 1 8 9 3
8 9 3 2 7 5 1 6 4
1 4 6 9 3 8 2 7 5
6 3 7 8 2 9 4 5 1
5 8 4 3 1 6 9 2 7
2 1 9 4 5 7 6 3 8
*/
