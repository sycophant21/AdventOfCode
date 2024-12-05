package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2024/Day5/H2/app/input.txt")
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
	var result uint64 = 0
	var invalidInputs = make([][]string, 0)
	positionMap := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else {
			input := strings.Split(line, "|")
			_, exists := positionMap[input[0]]
			if !exists {
				positionMap[input[0]] = make([]string, 0)
			}
			positionMap[input[0]] = append(positionMap[input[0]], input[1])
		}
	}
	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, ",")
		isInvalid := false
		for i := 0; i < len(inputs); i++ {
			for j := i + 1; j < len(inputs); j++ {
				res, err := search(positionMap, inputs[i], inputs[j])
				if err != nil || !res {
					isInvalid = true
					invalidInputs = append(invalidInputs, inputs)
					break
				}
			}
			if isInvalid {
				break
			}
		}
		isInvalid = false
	}
	validInputs := make([][]string, len(invalidInputs))
	for i, a := range invalidInputs {
		validInputs[i] = makeInputValid(a, positionMap)
		arr := validInputs[i]
		val, err := strconv.Atoi(arr[len(arr)/2])
		if err != nil {
			log.Fatal(err)
		} else {
			result += uint64(val)
		}
	}
	fmt.Println(result)
}

func search(input map[string][]string, key string, val string) (bool, error) {
	vals, exists := input[key]
	if exists {
		for _, v := range vals {
			if v == val {
				return true, nil
			}
		}
	} else {
		return false, errors.New("key doesn't exist")
	}
	return false, nil
}

func makeInputValid(input []string, inputMap map[string][]string) []string {
	validInput := make([]string, len(input))
	for i := 0; i < len(input); i++ {
		var newInput = make([]string, 0)
		newInput = append(newInput, input[0:i]...)
		newInput = append(newInput, input[i+1:]...)
		val := findMatches(newInput, inputMap[input[i]])
		validInput[len(input)-int(val)-1] = input[i]
	}
	return validInput
}
func findMatches(input []string, arr []string) uint16 {
	var result uint16 = 0
	for i := range arr {
		for j := range input {
			if arr[i] == input[j] {
				result++
			}
		}
	}
	return result
}
