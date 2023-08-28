package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

func part1() int {
	filePath := os.Args[1]
	// Read the input
	readFile, err := os.Open(filePath)

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var sum int = 0

	for fileScanner.Scan() {
		text := fileScanner.Text()
		p := strings.Split(text, " ")

		if p[0] == "A" { // Rock
			if p[1] == "X" { // Rock
				sum += 4
			}
			if p[1] == "Y" { // Paper
				sum += 8
			}
			if p[1] == "Z" { // Scissor
				sum += 3
			}
		}

		if p[0] == "B" { // Paper
			if p[1] == "X" { // Rock
				sum += 1
			}
			if p[1] == "Y" { // Paper
				sum += 5
			}
			if p[1] == "Z" { // Scissor
				sum += 9
			}
		}

		if p[0] == "C" { // Scissor
			if p[1] == "X" { // Rock
				sum += 7
			}
			if p[1] == "Y" { // Paper
				sum += 2
			}
			if p[1] == "Z" { // Scissor
				sum += 6
			}
		}
	}

	defer readFile.Close()

	return sum
}

func part2() int {
	filePath := os.Args[1]
	// Read the input
	readFile, err := os.Open(filePath)

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var sum int = 0

	for fileScanner.Scan() {
		text := fileScanner.Text()
		p := strings.Split(text, " ")

		if p[0] == "A" { // Rock
			if p[1] == "X" { // Lose
				sum += 3
			}
			if p[1] == "Y" { // Draw
				sum += 3 + 1
			}
			if p[1] == "Z" { // Win
				sum += 6 + 2
			}
		}

		if p[0] == "B" { // Paper
			if p[1] == "X" { // Lose
				sum += 1
			}
			if p[1] == "Y" { // Draw
				sum += 3 + 2
			}
			if p[1] == "Z" { // Win
				sum += 6 + 3
			}
		}

		if p[0] == "C" { // Scissor
			if p[1] == "X" { // Lose
				sum += 2
			}
			if p[1] == "Y" { // Draw
				sum += 3 + 3
			}
			if p[1] == "Z" { // Win
				sum += 6 + 1
			}
		}
	}

	defer readFile.Close()

	return sum
}

// Error checking
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
