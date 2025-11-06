package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createChessBoard(size uint64) string {
	const (
		blackTile = " "
		whiteTile = "#"
	)

	var boardString strings.Builder
	sizeInt := int(size)
	for i := 0; i < sizeInt*sizeInt; i++ {
		row := i / sizeInt
		col := i % sizeInt

		if (row+col)%2 == 0 {
			boardString.WriteString(blackTile)
		} else {
			boardString.WriteString(whiteTile)
		}

		if col == sizeInt-1 && row != sizeInt-1 {
			boardString.WriteString("\n")
		}
	}
	return boardString.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите желаемый размер шахматной доски: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	val, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		fmt.Println("Пожалуйста, введите целое неотрицательное число")
		return
	} else {
		fmt.Print(createChessBoard(val))
	}
}
