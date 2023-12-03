package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func isEnginePart(line string) bool {
	for _, char := range line {
		if char != '.' && !unicode.IsDigit(char) {
			return true
		}
	}

	return false
}

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

	result := 0
	matrix := []string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		matrix = append(matrix, line)
	}

	re := regexp.MustCompile(`\d+`)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if !unicode.IsDigit(rune(matrix[i][j])) {
				continue
			}

			num := re.FindString(matrix[i][j:])
			x, _ := strconv.Atoi(num)
			delta := len(num) - 1

			a := j
			if j > 0 {
				a = j - 1
			}

			b := j + delta
			if len(matrix[i])-j != delta+1 {
				b = j + delta + 2
			}

			if j > 0 && matrix[i][j-1] != '.' {
				result += x
			} else if len(matrix[i])-j != delta+1 && matrix[i][j+delta+1] != '.' {
				result += x
			} else if i > 0 && isEnginePart(matrix[i-1][a:b]) {
				result += x
			} else if i < len(matrix)-1 && isEnginePart(matrix[i+1][a:b]) {
				result += x
			}

			j += delta
		}
	}

	fmt.Println(result)
}
