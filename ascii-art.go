package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if os.Args[1] == "\\n" {
		return
	}
	to_print := os.Args[1]
	to_print_slice := strings.Split(to_print, "\\n")
	for i := 0; i < len(to_print_slice); i++ {
		//fmt.Println(to_print_slice[i])
		ascii_char := get_ascii_char(to_print_slice[i])
		show_ascii(ascii_char)
	}
	/* ascii_char := get_ascii_char(os.Args[1])
	show_ascii(ascii_char) */

}

func get_ascii_char(caractere string) []string {
	output := []string{}
	for _, element := range caractere {
		char := rune(element - 32)
		file, err := os.Open("standard.txt")
		if err != nil {
			os.Exit(1)
			fmt.Println(err)
		}
		scanner := bufio.NewScanner(file)
		ascii := []string{}
		for scanner.Scan() {
			ascii = append(ascii, scanner.Text())
		}
		for i := char * 9; i < (char*9)+9; i++ {
			output = append(output, ascii[i])
		}
	}
	return output
}
func show_ascii(ascii_char []string) {
	nbr_lettre := len(ascii_char) / 9
	var word_array []string
	for y := 1; y < 8; y++ {
		for i := y; i < nbr_lettre*9; i += 9 {
			word_array = append(word_array, ascii_char[i])
			if ascii_char[i] != "\n" {
				fmt.Print(ascii_char[i])
			}
		}
		fmt.Println("")
	}
}
