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
	var input = []int{}

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
		input = append(input, code)
	}

	n, v, err := getNounAndVerbFor(&input, 19690720)
	if err != nil {
		log.Fatal(err)
	}

	// Print result
	fmt.Println(100*n + v)
}

func getNounAndVerbFor(in *[]int, n int) (noun int, verb int, err error) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			result, err := Execute(in, i, j)
			if err != nil {
				log.Fatal(err)
			}
			if result[0] == n {
				return i, j, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("No solution found")
}

// Execute : runs Intcode
func Execute(in *[]int, noun int, verb int) ([]int, error) {
	// Make a new slice
	op := make([]int, len(*in))
	copy(op, *in)

	op[1] = noun
	op[2] = verb
	for i := 0; i < len(op); i++ {
		switch op[i] {
		case 99:
			return op, nil
		case 1:
			op[op[i+3]] = op[op[i+1]] + op[op[i+2]]
			i += 3
		case 2:
			op[op[i+3]] = op[op[i+1]] * op[op[i+2]]
			i += 3
		}
	}
	return nil, fmt.Errorf("Program exited unexpectedly")
}
