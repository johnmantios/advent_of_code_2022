package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func solution_first_part(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	hash_map := make(map[int]float64)
	counter := 0

	for scanner.Scan() { // internally, it advances token based on seperator
		if scanner.Text() != "" {
			value, err := strconv.ParseFloat(scanner.Text(), 2)
			if err != nil {
				log.Fatal(err)
			}
			hash_map[counter] += value
		} else {
			counter++
		}
	}

	solution := math.Inf(-1)

	for _, value := range hash_map {
		if value > solution {
			solution = value
		}
	}
	fmt.Println(solution)
}

func solution_second_part(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	hash_map := make(map[int]float64)
	counter := 0

	for scanner.Scan() { // internally, it advances token based on seperator
		if scanner.Text() != "" {
			value, err := strconv.ParseFloat(scanner.Text(), 2)
			if err != nil {
				log.Fatal(err)
			}
			hash_map[counter] += value
		} else {
			counter++
		}
	}

	largest_three := []float64{math.Inf(-1), math.Inf(-1), math.Inf(-1)}

	for _, value := range hash_map {
		if value > largest_three[0] {
			largest_three[0] = value
			sort.Float64s(largest_three)
		}
	}

	var solution float64

	for _, value := range largest_three {
		solution += value
	}

	fmt.Println(solution)
}

func main() {
	solution_first_part("input1.txt")
	solution_second_part("input1.txt")
}
