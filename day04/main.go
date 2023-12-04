package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Scratchcard struct {
	winningNumbers []string
	myNumbers      []string
}

func numberOfGames(scratchcards map[int]Scratchcard, i int, result *int) {
	mb := make(map[string]struct{}, len(scratchcards[i].myNumbers))
	for _, x := range scratchcards[i].myNumbers {
		mb[x] = struct{}{}
	}

	j := 0
	for _, x := range scratchcards[i].winningNumbers {
		if _, found := mb[x]; found {
			*result++
			j++
			numberOfGames(scratchcards, i+j, result)
		}
	}
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

	re1 := regexp.MustCompile(`Card\s+\d+: `)
	re2 := regexp.MustCompile(`\s+`)
	re3 := regexp.MustCompile(`\d+`)
	result := 0
	scratchcards := make(map[int]Scratchcard)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		x := re3.FindString(line)
		i, _ := strconv.Atoi(x)
		line = re1.ReplaceAllString(line, "${1}")
		scratchcard := strings.Split(line, " | ")
		winningNumbers := re2.Split(scratchcard[0], -1)
		myNumbers := re2.Split(scratchcard[1], -1)
		scratchcards[i] = Scratchcard{winningNumbers, myNumbers}
		result++
	}

	for i := 1; i <= len(scratchcards); i++ {
		numberOfGames(scratchcards, i, &result)
	}

	fmt.Println(result)
}
