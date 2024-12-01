package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("2024/Day1/H1/app/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		er := file.Close()
		if er != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	var result uint64
	var one []uint64
	var two []uint64
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
					one = append(one, last)
					last = 0
				}
			} else {
				if b >= '0' && b <= '9' {
					last = (last * uint64(10)) + uint64(b-48)
					if i == len(line)-1 {
						two = append(two, last)
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
	sort.Slice(one, func(x, y int) bool {
		return one[x] < one[y]
	})
	sort.Slice(two, func(x, y int) bool {
		return two[x] < two[y]
	})
	for i := 0; i < len(one); i++ {
		var res = int64(one[i] - two[i])
		if res < 0 {
			result += uint64(-res)
		} else {
			result += uint64(res)
		}
	}
	fmt.Print(result)
}
