package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]",

			"EX: go run . something standard")
	} else {
		lines, err := linesFromReader(os.Args[2:3][0] + ".txt")
		check(err)
		arr := strings.Split(os.Args[1:2][0], "\\n")

		printASCII(arr, lines)
	}
}

func printASCII(arr, lines []string) {
	for _, e := range arr {
		for i := 0; i < 8; i++ {
			printS := ""
			r := []rune(e)
			for _, runeR := range r {
				printS += lines[int(runeR-32)*9+1+i]
			}
			if printS != "" {
				fmt.Println(printS)
			}
		}
	}
}

func linesFromReader(filename string) ([]string, error) {
	f, _ := os.Open(filename) // opens file
	defer f.Close() 

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
