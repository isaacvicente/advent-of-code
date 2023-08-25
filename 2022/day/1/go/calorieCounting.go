package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filePath := os.Args[1]
	// Read the input
	readFile, err := os.Open(filePath)

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		// Append all the lines in the string array
		fileLines = append(fileLines, fileScanner.Text())
	}

	defer readFile.Close()

	var sum int = 0
	var calories int = 0
	for _, line := range fileLines {
		if line != "" {
			// Convert string to integer
			num, err := strconv.Atoi(line)
			check(err)
			sum += num
		} else {
			if sum > calories {
				calories = sum
			}
			// Reset the sum for the next elf
			sum = 0
		}
	}

	fmt.Println(calories)
}

// Error checking
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
