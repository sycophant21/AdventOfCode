package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2024/Day7/H1/app/input_temp.txt")
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
	var result uint64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ":")
		res, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		params := ExtractParams(strings.Split(split[1], " "))
		if Check(uint64(res), params) {
			result += uint64(res)
		}
	}
	fmt.Println(result)
}

func Check(res uint64, params []uint64) bool {
	resTemp := []uint64{params[0]}
	ops := []string{"+", "*"}
	for i := 1; i < len(params); i++ {
		resTemp = Calc(ops, resTemp, params[i])
	}
	for _, r := range resTemp {
		if r == res {
			return true
		}
	}
	return false
}

func Calc(ops []string, nums []uint64, num uint64) []uint64 {
	res := make([]uint64, 0)
	for _, val := range nums {
		for _, op := range ops {
			var n uint64
			if op == "+" {
				n = num + val
			} else {
				n = num * val
			}
			res = append(res, n)
		}
	}
	return res
}

func ExtractParams(params []string) []uint64 {
	res := make([]uint64, 0)
	for _, s := range params {
		if s != "" {
			r, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
				return nil
			}
			res = append(res, uint64(r))
		}
	}
	return res
}
