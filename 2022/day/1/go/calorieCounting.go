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
	readFile, err := os.Open(filePath)

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var sum int = 0
	var calories int = 0
	for _, line := range fileLines {
		if line != "" {
			num, err := strconv.Atoi(line)
			check(err)
			sum += num
		} else {
			if sum > calories {
				calories = sum
			}
			sum = 0
		}
	}

	fmt.Println(calories)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
