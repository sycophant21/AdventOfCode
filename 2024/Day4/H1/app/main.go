package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := "2024/Day4/H1/app/input.txt"
	file, err := os.Open(fileName)
	defer func(file *os.File) {
		er := file.Close()
		if er != nil {
			log.Fatal(er)
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
			if s == "X" {
				if findRune(&input, "M", []int{i - 1, j}) {
					if findRune(&input, "A", []int{i - 2, j}) {
						if findRune(&input, "S", []int{i - 3, j}) {
							fmt.Println("x:", i, "y:", j, "s:", s, "branch: 1")
							result++
						}
					}
				}
				if findRune(&input, "M", []int{i + 1, j}) {
					if findRune(&input, "A", []int{i + 2, j}) {
						if findRune(&input, "S", []int{i + 3, j}) {
							fmt.Println("x:", i, "y:", j, "s:", s, "branch: 2")
							result++
						}
					}
				}
				if findRune(&input, "M", []int{i - 1, j - 1}) {
					if findRune(&input, "A", []int{i - 2, j - 2}) {
						if findRune(&input, "S", []int{i - 3, j - 3}) {
							fmt.Println("x:", i, "y:", j, "s:", s, "branch: 3")
							result++
						}
					}
				}
				if findRune(&input, "M", []int{i + 1, j - 1}) {
					if findRune(&input, "A", []int{i + 2, j - 2}) {
						if findRune(&input, "S", []int{i + 3, j - 3}) {
							fmt.Println("x:", i, "y:", j, "s:", s, "branch: 4")
							result++
						}
					}
				}
				if findRune(&input, "M", []int{i - 1, j + 1}) {
					if findRune(&input, "A", []int{i - 2, j + 2}) {
						if findRune(&input, "S", []int{i - 3, j + 3}) {
							fmt.Println("x:", i, "y:", j, "s:", s, "branch: 5")
							result++
						}
					}
				}
				if findRune(&input, "M", []int{i + 1, j + 1}) {
					if findRune(&input, "A", []int{i + 2, j + 2}) {
						if findRune(&input, "S", []int{i + 3, j + 3}) {
							fmt.Println("x:", i, "y:", j, "s:", s, "branch: 6")
							result++
						}
					}
				}
				if findRune(&input, "M", []int{i, j - 1}) {
					if findRune(&input, "A", []int{i, j - 2}) {
						if findRune(&input, "S", []int{i, j - 3}) {
							fmt.Println("x:", i, "y:", j, "s:", s, "branch: 7")
							result++
						}
					}
				}
				if findRune(&input, "M", []int{i, j + 1}) {
					if findRune(&input, "A", []int{i, j + 2}) {
						if findRune(&input, "S", []int{i, j + 3}) {
							fmt.Println("x:", i, "y:", j, "s:", s, "branch: 8")
							result++
						}
					}
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
