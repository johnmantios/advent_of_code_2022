package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func solution_first_part(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	my_points := map[string]int{"X": 1, "Y": 2, "Z": 3}
	score := 0

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")

		if data[1] == "X" && data[0] == "A" {
			score += my_points[data[1]] + 3
		} else if data[1] == "X" && data[0] == "B" {
			score += my_points[data[1]]
		} else if data[1] == "X" && data[0] == "C" {
			score += my_points[data[1]] + 6
		} else if data[1] == "Y" && data[0] == "A" {
			score += my_points[data[1]] + 6
		} else if data[1] == "Y" && data[0] == "B" {
			score += my_points[data[1]] + 3
		} else if data[1] == "Y" && data[0] == "C" {
			score += my_points[data[1]]
		} else if data[1] == "Z" && data[0] == "A" {
			score += my_points[data[1]]
		} else if data[1] == "Z" && data[0] == "B" {
			score += my_points[data[1]] + 6
		} else if data[1] == "Z" && data[0] == "C" {
			score += my_points[data[1]] + 3
		}
	}

	fmt.Println(score)

}

func solution_second_part(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	my_points := map[string]int{"X": 1, "Y": 2, "Z": 3}
	lose := map[string]string{"A": "Z", "B": "X", "C": "Y"}
	tie := map[string]string{"A": "X", "B": "Y", "C": "Z"}
	win := map[string]string{"A": "Y", "B": "Z", "C": "X"}
	score := 0

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")

		if data[1] == "X" {
			score += my_points[lose[data[0]]]
		} else if data[1] == "Y" {
			score += my_points[tie[data[0]]] + 3
		} else if data[1] == "Z" {
			score += my_points[win[data[0]]] + 6
		}
	}

	fmt.Println(score)

}

func main() {
	solution_first_part("input2.txt")
	solution_second_part("input2.txt")
}
