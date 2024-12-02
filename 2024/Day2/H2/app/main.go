package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
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
		line := scanner.Text() + " "
		vals := strings.Split(line, " ")
		var arr []uint8
		for _, s := range vals {
			val, err := strconv.Atoi(s)
			if err == nil {
				arr = append(arr, uint8(val))
			}
		}
		var isInvalid = true
		for i := range arr {
			var arrTemp = make([]uint8, 0)
			arrTemp = append(arrTemp, arr[0:i]...)
			arrTemp = append(arrTemp, arr[i+1:]...)
			res := isValid(arrTemp)
			if res {
				isInvalid = false
				break
			}

		}
		if !isInvalid {
			result++
		}
	}
	fmt.Println(result)
}

func isValid(arr []uint8) bool {
	var arrInc = make([]uint8, 0)
	var arrDec = make([]uint8, 0)
	for _, val := range arr {
		arrInc = append(arrInc, val)
		arrDec = append(arrDec, val)
	}
	sort.Slice(arrInc, func(x, y int) bool {
		return arrInc[x] < arrInc[y]
	})
	sort.Slice(arrDec, func(x, y int) bool {
		return arrDec[x] > arrDec[y]
	})
	return (reflect.DeepEqual(arr, arrInc) && isSliceValid(arr, true)) || (reflect.DeepEqual(arr, arrDec) && isSliceValid(arr, false))
}

func isSliceValid(arr []uint8, inc bool) bool {
	for i := 0; i < len(arr); i++ {
		if i > 0 {
			res := int8(arr[i] - arr[i-1])
			if res == 0 {
				return false
			} else {
				if (res > 0 != inc) || (res > 3 || res < -3) {
					return false
				}
			}
		}
	}
	return true
}
