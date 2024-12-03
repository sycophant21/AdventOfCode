package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("2024/Day3/H2/app/input.txt")
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
	var do = true
	for scanner.Scan() {
		line := scanner.Text()
		for i, r := range line {
			if r == 'd' && i < len(line)-1 {
				o := line[i+1 : i+2]
				if i < len(line)-4 {
					n := line[i+2 : i+3]
					a := line[i+3 : i+4]
					t := line[i+4 : i+5]
					if o == "o" {
						if n == "n" && a == "'" && t == "t" {
							do = false
						} else {
							do = true
						}
					}
				} else {
					do = true
				}
			}
			if r == 'm' && i < len(line)-3 && do {
				u := line[i+1 : i+2]
				l := line[i+2 : i+3]
				ob := line[i+3 : i+4]
				isValid := true
				if u == "u" && l == "l" && ob == "(" {
					var n1 uint64
					var n2 uint64
					var index = i + 4
					for j := index; j < len(line); j++ {
						var s = []rune(line[j : j+1])
						if s[0] == ',' {
							index = j + 1
							isValid = true
							break
						} else {
							if s[0] >= '0' && s[0] <= '9' {
								n1 = (n1 * 10) + uint64(s[0]-48)
							} else {
								isValid = false
								break
							}
						}
					}
					if isValid {
						isValid = false
						for j := index; j < len(line); j++ {
							var s = []rune(line[j : j+1])
							if s[0] == ')' {
								isValid = true
								break
							} else {
								if s[0] >= '0' && s[0] <= '9' {
									n2 = (n2 * 10) + uint64(s[0]-48)
								} else {
									isValid = false
									break
								}
							}
						}
					}
					if isValid {
						result += n1 * n2
					}
				}
			}
		}
	}
	fmt.Println(result)
}
