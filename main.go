package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// fondid := []string{"standard.txt", "shadow.txt", "thinkertoy.txt"}

	fontType := os.Args[2]
	fontTypeFile := fontType + ".txt"
	fmt.Println(fontTypeFile)

	font, err := ioutil.ReadFile(fontTypeFile)
	if err != nil {
		fmt.Println("Cant find file")
	}
	art := strings.Split(string(font), "\n")
	input := []byte(os.Args[1])

	var inputbyuser []int // user inputi teha byte kujule
	for _, index := range input {
		inputbyuser = append(inputbyuser, int(index))
	}

	for x := 1; x < 9; x++ {
		output := ""
		for ch := range inputbyuser {
			output = output + art[((inputbyuser[ch]-32)*9)+x]
		}
		fmt.Println(output)
	}
}
