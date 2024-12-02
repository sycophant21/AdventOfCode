package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("2024/Day2/H1/app/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		er := file.Close()
		if er != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	var result uint16
	for scanner.Scan() {
		line := scanner.Text()
		isInc := false
		isDec := false
		var lastVal uint8
		var val uint8
		isInvalid := false
		for _, r := range line {
			if r >= '0' && r <= '9' {
				val = (val * 10) + uint8(r-48)
			} else {
				if lastVal == 0 {
					lastVal = val
				} else {
					if lastVal < val && val-lastVal <= 3 {
						if !isInc && !isDec {
							isInc = true
						} else {
							if isDec {
								isInvalid = true
								break
							}
						}
					} else if lastVal > val && lastVal-val <= 3 {
						if !isInc && !isDec {
							isDec = true
						} else {
							if isInc {
								isInvalid = true
								break
							}
						}
					} else {
						isInvalid = true
						break
					}
				}
				lastVal = val
				val = 0
			}
		}
		if !isInvalid {
			if (lastVal < val && val-lastVal <= 3 && isInc) || (lastVal > val && lastVal-val <= 3 && isDec) {
				fmt.Println(line)
				result++
			}
			isInvalid = false
		}
	}
	fmt.Println(result)
}
