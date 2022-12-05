package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var priorities = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

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
	var group []string
	score := 0

	for scanner.Scan() {
		line := scanner.Text()

		group = append(group, line[:len(line)/2])
		group = append(group, line[len(line)/2:])

		if len(group) == 2 {
			for _, char := range group[0] {
				if strings.Contains(group[1], string(char)) {
					score += priorities[string(char)]
					group = nil
					break
				}
			}
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
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	var group []string
	score := 0

	for scanner.Scan() {
		line := scanner.Text()

		group = append(group, line)

		if len(group) == 3 {
			for _, char := range group[0] {
				if strings.Contains(group[1], string(char)) && strings.Contains(group[2], string(char)) {
					score += priorities[string(char)]
					group = nil
					break
				}
			}
		}

	}
	fmt.Println(score)
}

func main() {
	solution_first_part("input3.txt")
	solution_second_part("input3.txt")
}
