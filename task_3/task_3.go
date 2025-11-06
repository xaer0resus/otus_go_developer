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

	var boardBuilder strings.Builder
	sizeInt := int(size)
	for i := 0; i < sizeInt; i++ {
		var pattern string
		if i%2 == 0 {
			pattern = blackTile + whiteTile
		} else {
			pattern = whiteTile + blackTile
		}
		boardBuilder.WriteString(strings.Repeat(pattern, sizeInt/2) + "\n")
	}
	return boardBuilder.String()
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
