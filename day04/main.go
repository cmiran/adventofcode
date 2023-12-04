package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file path")
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	re1 := regexp.MustCompile(`Card\s+\d+: `)
	re2 := regexp.MustCompile(`\s+`)
	result := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = re1.ReplaceAllString(line, "${1}")
		scratchcard := strings.Split(line, " | ")
		winningNumbers := re2.Split(scratchcard[0], -1)
		myNumbers := re2.Split(scratchcard[1], -1)

		mb := make(map[string]struct{}, len(myNumbers))
		for _, x := range myNumbers {
			mb[x] = struct{}{}
		}

		points := 0
		for _, x := range winningNumbers {
			if _, found := mb[x]; found {
				if points == 0 {
					points++
				} else {
					points *= 2
				}
			}
		}

		result += points
	}

	fmt.Println(result)
}
