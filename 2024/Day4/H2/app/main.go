package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2024/Day4/H2/app/input.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	input := make([][]string, 0)
	var result uint64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, strings.Split(line, ""))
	}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			s := input[i][j]
			if s == "A" {
				m11 := []int{i - 1, j - 1}
				m12 := []int{i + 1, j + 1}
				m21 := []int{i + 1, j - 1}
				m22 := []int{i - 1, j + 1}
				s11 := []int{i + 1, j + 1}
				s12 := []int{i - 1, j - 1}
				s21 := []int{i - 1, j + 1}
				s22 := []int{i + 1, j - 1}
				if ((findRune(&input, "M", m11) && findRune(&input, "S", s11)) || (findRune(&input, "M", m12) && findRune(&input, "S", s12))) && ((findRune(&input, "M", m21) && findRune(&input, "S", s21)) || (findRune(&input, "M", m22) && findRune(&input, "S", s22))) {
					result++
				}
			}
		}
	}
	fmt.Println(result)
}

func findRune(input *[][]string, s string, index []int) bool {
	x := index[0]
	y := index[1]
	if (x >= 0 && x <= len(*input)-1) && (y >= 0 && y <= len((*(input))[x])-1) {
		return (*input)[x][y] == s
	} else {
		return false
	}
}
