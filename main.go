package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Введите целое число: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	s := scanner.Text()
	fmt.Printf("Вы ввели число: %v\n", s)
}
