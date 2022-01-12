package main

import "fmt"

func main() {
	fmt.Println("start")
	str1 := "привет"
	fmt.Println(str1[1:4])
	fmt.Println(len(str1))
	fmt.Println(str1)

	runes := []rune(str1)
	fmt.Println(string(runes[0]))
	runes[1] = 'с'
	fmt.Println(string(runes))
}
