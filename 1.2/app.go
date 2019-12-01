package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var err error
	var fuelRequired = 0

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fuelRequired += calculateFuelForModule(mass)
	}

	fmt.Printf("Total fuel required: %d\n", fuelRequired)
}

func calculateFuelForModule(m int) int {
	if fuel := m/3 - 2; fuel > 0 {
		return fuel + calculateFuelForModule(fuel)
	}
	return 0
}
