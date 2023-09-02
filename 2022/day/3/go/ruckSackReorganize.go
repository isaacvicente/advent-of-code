package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var repeated []rune
	var sum int

	for scanner.Scan() {
		input := scanner.Text()
		middle := len(input) / 2
		left := []rune(input[:middle])
		right := []rune(input[middle:])

		for i := 0; i < len(left); i++ {
			char := search(right, left[i])
			if char != '0' {
				repeated = append(repeated, char)
				sum += priority(char)
				break
			}
		}
	}

	fmt.Println("Part 1 - Sum of priorities:", sum)
}

func part2() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var repeated []rune
	var common []rune
	var sum int
	var lines []string
	var first []rune
	var second []rune
	var third []rune

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i += 3 {
		first = []rune(lines[i])
		second = []rune(lines[i+1])
		third = []rune(lines[i+2])

		for i := 0; i < len(first); i++ {
			char := search(second, first[i])
			if char != '0' {
				repeated = append(repeated, char)
			}
		}

		for i := len(repeated) - 1; i >= 0; i-- {
			char := search(third, repeated[i])
			if char != '0' {
				common = append(common, char)
				break
			}
		}
	}

	for j := 0; j < len(common); j++ {
		sum += priority(common[j])
	}

	fmt.Println("Part 2 - Sum of priorities:", sum)
}

func priority(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - 38
	} else {
		return int(r) - 96
	}
}

func search(a []rune, k rune) rune {
	for i := 0; i < len(a); i++ {
		if a[i] == k {
			return k
		}
	}
	return '0'
}
