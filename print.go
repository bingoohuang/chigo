package chigo

import (
	"bufio"
	"fmt"
	"strings"
)

func PrintWithColors(text string) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	for i := 1.0; scanner.Scan(); i++ {
		rgb := NewRGB(i)
		fmt.Printf("%sm\n", rgb.Color(scanner.Text()))
	}
}
