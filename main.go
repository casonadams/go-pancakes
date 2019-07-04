package main

import (
	"fmt"

	waiter "github.com/casonadams/go-pancakes/waiter"
)

func main() {
	input := []string{"-", "+", "-", "-", "-", "+", "-", "+", "+", "-", "-", "+", "+"}

	guy := waiter.NewWaiter()
	output, err := guy.Organize(input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Number of flips:\t%v\n", output)
}
