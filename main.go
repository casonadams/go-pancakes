package main

import (
	"fmt"

	waiter "github.com/casonadams/go-pancakes/waiter"
)

// TODO: add features to take a string or read from a file from commandline
func main() {
	var inputtest = []struct {
		input  string
		output int
	}{
		{"-", 1},    // Case #1
		{"-+", 1},   // Case #2
		{"+-", 2},   // Case #3
		{"+++", 0},  // Case #4
		{"--+-", 3}, // Case #5
	}
	w := waiter.NewWaiter()
	fmt.Println("Output\t\tInput")

	for i, tt := range inputtest {
		expected := tt.output
		actual, _, err := w.Organize(tt.input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			if actual == expected {
				fmt.Printf("Case #%v: %v\t%v\n", i+1, actual, tt.input)
			}
		}
	}
}
