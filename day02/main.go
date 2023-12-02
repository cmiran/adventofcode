package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Set struct {
	Red   int
	Blue  int
	Green int
}

func splitLine(line, r string) []string {
	re := regexp.MustCompile(r)
	return re.Split(line, -1)
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

	for fileScanner.Scan() {
		line := fileScanner.Text()[5:]

		content := splitLine(line, ": ")
		sets := splitLine(content[1], "; ")
		s1 := Set{}

		for _, set := range sets {
			cubes := splitLine(set, ", ")
			s2 := Set{}

			for _, cube := range cubes {
				content := splitLine(cube, " ")
				x, _ := strconv.Atoi(content[0])

				switch content[1] {
				case "red":
					if s2.Red += x; s2.Red > s1.Red {
						s1.Red = s2.Red
					}
				case "blue":
					if s2.Blue += x; s2.Blue > s1.Blue {
						s1.Blue = s2.Blue
					}
				case "green":
					if s2.Green += x; s2.Green > s1.Green {
						s1.Green = s2.Green
					}
				}
			}
		}

    result += s1.Red * s1.Blue * s1.Green
	}

	fmt.Println(result)
}
