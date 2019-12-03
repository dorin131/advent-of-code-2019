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
	var str string
	var codes = []int{}

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	// Read input as text
	for scanner.Scan() {
		if len(str) > 0 && str[len(str)-1] != ',' {
			str = str + ","
		}
		str = str + scanner.Text()
	}

	// Convert text input to slice of integers
	for _, v := range strings.Split(str, ",") {
		code, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		codes = append(codes, code)
	}

	result, err := Execute(codes)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range result {
		if i == len(result)-1 {
			fmt.Printf("%d\n", v)
			continue
		}
		fmt.Printf("%d,", v)
	}
}

// Execute : runs Intcode
func Execute(codes []int) ([]int, error) {
	for i := 0; i < len(codes); i++ {
		switch codes[i] {
		case 99:
			return codes, nil
		case 1:
			codes[codes[i+3]] = codes[codes[i+1]] + codes[codes[i+2]]
			i += 3
		case 2:
			codes[codes[i+3]] = codes[codes[i+1]] * codes[codes[i+2]]
			i += 3
		}
	}
	return nil, fmt.Errorf("Program exited unexpectedly")
}
