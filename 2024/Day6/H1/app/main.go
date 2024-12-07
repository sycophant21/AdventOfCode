package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("2024/Day6/H1/app/input.txt")
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
	var direction = UP
	mat := make([][]byte, 0)
	currPos := make([]uint16, 0)
	out := false
	rowNo := uint16(0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]byte, 0)
		for i := range line {
			s := line[i : i+1]
			appending := uint8(2)
			if s == "." {
				appending = 0
			} else if s == "#" {
				appending = 1
			} else {
				currPos = append(currPos, rowNo)
				currPos = append(currPos, uint16(i))
			}
			row = append(row, appending)
		}
		mat = append(mat, row)
		rowNo++
	}
	for {
		mat[currPos[0]][currPos[1]] = 3
		currPos, direction, out = GetNextPos(direction, currPos, mat)
		if out {
			break
		}
	}
	var result uint64 = 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 3 {
				result++
			}
		}
	}
	fmt.Println(result)
}

func GetNextPos(dir Direction, currPos []uint16, mat [][]uint8) ([]uint16, Direction, bool) {
	isOut := IsAtEdge(dir, currPos, mat)
	if isOut {
		return nil, dir, true
	} else {
		newArr := make([]uint16, 0)
		if dir == UP {
			if mat[currPos[0]-1][currPos[1]] == 1 {
				return currPos, RIGHT, false
			} else {
				newArr = append(newArr, currPos[0]-1)
				newArr = append(newArr, currPos[1])
			}
		} else if dir == RIGHT {
			if mat[currPos[0]][currPos[1]+1] == 1 {
				return currPos, DOWN, false
			} else {
				newArr = append(newArr, currPos[0])
				newArr = append(newArr, currPos[1]+1)
			}
		} else if dir == DOWN {
			if mat[currPos[0]+1][currPos[1]] == 1 {
				return currPos, LEFT, false
			} else {
				newArr = append(newArr, currPos[0]+1)
				newArr = append(newArr, currPos[1])
			}
		} else {
			if mat[currPos[0]][currPos[1]-1] == 1 {
				return currPos, UP, false
			} else {
				newArr = append(newArr, currPos[0])
				newArr = append(newArr, currPos[1]-1)
			}
		}
		return newArr, dir, false
	}
}

func IsAtEdge(dir Direction, currPos []uint16, mat [][]uint8) bool {
	if dir == UP {
		return currPos[0] == 0
	} else if dir == RIGHT {
		return currPos[1] == uint16(len(mat[currPos[0]]))-1
	} else if dir == DOWN {
		return currPos[0] == uint16(len(mat))-1
	} else {
		return currPos[1] == 0
	}
}

type Direction uint8

const (
	UP    Direction = 0
	RIGHT Direction = 1
	DOWN  Direction = 2
	LEFT  Direction = 3
)
