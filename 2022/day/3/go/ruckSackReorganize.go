package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
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
				if unicode.IsUpper(char) {
					sum += int(char) - 38
				} else {
					sum += int(char) - 96
				}
				break
			}
		}
	}

	fmt.Println("Sum of priorities:", sum)
}

func search(a []rune, k rune) rune {
	for i := 0; i < len(a); i++ {
		if a[i] == k {
			return k
		}
	}
	return '0'
}
