package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("2024/Day1/H2/app/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		er := file.Close()
		if er != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	var result = make(map[uint64][]uint8)
	for scanner.Scan() {
		line := scanner.Text()
		var readingOne = true
		var readingTwo = false
		var last uint64
		for i, b := range line {
			if readingOne {
				if b >= '0' && b <= '9' {
					last = (last * uint64(10)) + uint64(b-48)
				} else {
					readingOne = false
					_, exist := result[last]
					if exist {
						result[last][0]++
					} else {
						result[last] = append(append(make([]uint8, 0), 1), 0)
					}
					last = 0
				}
			} else {
				if b >= '0' && b <= '9' {
					last = (last * uint64(10)) + uint64(b-48)
					if i == len(line)-1 {
						_, exist := result[last]
						if exist {
							result[last][1]++
						} else {
							result[last] = append(append(make([]uint8, 0), 0), 1)
						}
						break
					}
					if !readingTwo {
						readingTwo = true
					}
				} else {
					continue
				}
			}
		}
	}
	var count uint64
	for a, b := range result {
		count += a * uint64(b[0]) * uint64(b[1])
	}
	fmt.Print(count)
}
