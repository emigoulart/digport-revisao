package main

import "fmt"

func main() {

	dezenas := []int{10, 20, 40, 30, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(dezenas), cap(dezenas), dezenas)

	fmt.Printf("len=%d cap=%d %v\n", len(dezenas[:0]), cap(dezenas[:0]), dezenas[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(dezenas[2:]), cap(dezenas[2:]), dezenas[2:])

	dezenas = append(dezenas, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(dezenas[:2]), cap(dezenas[:2]), dezenas[:2])

	fmt.Printf("len=%d cap=%d %v\n", len(dezenas), cap(dezenas), dezenas)

	test := []int{}
	fmt.Printf("cap=%d\n", cap(test))
	test = append(test, 1)
	fmt.Printf("cap=%d\n", cap(test))
}
