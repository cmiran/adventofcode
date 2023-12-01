package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
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

	calibrationValue := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		length := len(line)
		digit := ""

		for i := 0; i < length; i++ {
			c := rune(line[i])
			if unicode.IsDigit(c) {
				digit += string(c)

				break
			}
		}

		for i := length - 1; i >= 0; i-- {
			c := rune(line[i])
			if unicode.IsDigit(c) {
				digit += string(c)

				break
			}
		}

		if x, err := strconv.Atoi(digit); err == nil {
			calibrationValue += x
		}
	}

  fmt.Println("Calibration value: ", calibrationValue)
}
