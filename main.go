package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input data: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	s := scanner.Text()
	fmt.Printf("You entered data: %v\n", s)
}
