package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
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

		lines, err := urlToLines("https://git.01.kood.tech/root/public/raw/branch/master/subjects/ascii-art/" + os.Args[2:3][0] + ".txt")
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

func urlToLines(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return linesFromReader(resp.Body)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
