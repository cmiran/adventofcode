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

var re = regexp.MustCompile(`\d+`)

func findEnginePart(line string, index int) []int {
	gears := []int{}
	idxs := re.FindAllStringIndex(line, -1)

	for _, idx := range idxs {
		if (index >= idx[0] && index <= idx[1]) || index+1 == idx[0] {
			g, _ := strconv.Atoi(line[idx[0]:idx[1]])
			gears = append(gears, g)
		}
	}

	return gears
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

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != '*' {
				continue
			}

			gears := []int{}

			if j > 0 && unicode.IsDigit(rune(matrix[i][j-1])) {
				nums := re.FindAllString(matrix[i][:j], -1)
				g, _ := strconv.Atoi(nums[len(nums)-1])
				gears = append(gears, g)
			}

			if j < len(matrix[i])-1 && unicode.IsDigit(rune(matrix[i][j+1])) {
				num := re.FindString(matrix[i][j+1:])
				g, _ := strconv.Atoi(num)
				gears = append(gears, g)
			}

			if i > 0 {
				g := findEnginePart(matrix[i-1], j)
				if len(g) > 0 {
					gears = append(gears, g...)
				}
			}

			if i < len(matrix)-1 {
				g := findEnginePart(matrix[i+1], j)
				if len(g) > 0 {
					gears = append(gears, g...)
				}
			}

			if len(gears) > 1 {
				result += gears[0] * gears[1]
			}
		}
	}

	fmt.Println(result)
}
