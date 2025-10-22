package main

import (
	"fmt"
	"strings"
)

func createChessBoard(size int) string {
	const (
		blackTile = " "
		whiteTile = "#"
	)

	var boardBuilder strings.Builder
	for i := 0; i < size; i++ {
		var pattern string
		if i%2 == 0 {
			pattern = blackTile + whiteTile
		} else {
			pattern = whiteTile + blackTile
		}
		boardBuilder.WriteString(strings.Repeat(pattern, size/2) + "\n")
	}
	return boardBuilder.String()
}

func main() {
	var size int
	fmt.Print("Введите желаемый размер шахматной доски: ")
	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println("Ошибка ввода: введите целое число")
		return
	}
	if size <= 0 {
		fmt.Println("Ошибка ввода: введите положительное число")
	}
	fmt.Print(createChessBoard(size))
}
