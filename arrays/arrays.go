package main

import "fmt"

func displayBoard(board [8][8]rune) {
	columns := [8]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	fmt.Printf("%2c", ' ')
	for _, column := range columns {
		fmt.Printf("%2c", column)
	}
	fmt.Println()
	for rindex, row := range board {
		fmt.Printf("%2v", 8-rindex)
		for _, column := range row {
			fmt.Printf("%2c", column)

		}
		fmt.Printf("%2v", 8-rindex)
		fmt.Println()
	}
	fmt.Printf("%2c", ' ')
	for _, column := range columns {
		fmt.Printf("%2c", column)
	}
	fmt.Println()
}

func main() {

	var board [8][8]rune

	black := "kqrbnp"
	white := "KQRBNP"

	board[0][0] = rune(black[2])
	board[0][1] = rune(black[4])
	board[0][2] = rune(black[3])
	board[0][3] = rune(black[1])
	board[0][4] = rune(black[0])
	board[0][5] = rune(black[3])
	board[0][6] = rune(black[4])
	board[0][7] = rune(black[2])
	board[7][0] = rune(white[2])
	board[7][1] = rune(white[4])
	board[7][2] = rune(white[3])
	board[7][3] = rune(white[1])
	board[7][4] = rune(white[0])
	board[7][5] = rune(white[3])
	board[7][6] = rune(white[4])
	board[7][7] = rune(white[2])

	for column := range board[1] {
		board[1][column] = rune(black[5])
		board[2][column] = ' '
		board[3][column] = ' '
		board[4][column] = ' '
		board[5][column] = ' '
		board[6][column] = rune(white[5])
	}

	displayBoard(board)

}
