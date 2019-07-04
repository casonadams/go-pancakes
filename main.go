package main

import (
	"fmt"

	waiter "github.com/casonadams/go-pancakes/waiter"
)

func main() {
	input := []string{"-", "+", "-", "-", "-", "+", "-", "+", "+", "-", "-", "+", "+"}

	w := waiter.NewWaiter()
	flipCount, _, err := w.Organize(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of flips:\t%v\n", flipCount)
}
