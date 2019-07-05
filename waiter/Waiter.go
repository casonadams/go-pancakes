package waiter

import (
	"errors"
	"fmt"
	"strings"
)

// Waiter allows access to the flip count
type Waiter struct {
	flipCount int
}

// NewWaiter Creates and Returns a new Waiter
func NewWaiter() *Waiter {
	return &Waiter{}
}

// Organize flips all pancakes to be happy face up adhearing to physical bounds of pancake flipping
func (w *Waiter) Organize(s string) (int, []string, error) {
	w.flipCount = 0
	var err error

	// Convert string to slice for convience
	arry := strings.Split(s, "")

	// Make sure the data is handled right to left. (backwards to how I think about it)
	arry = w.reverse(arry)

	// Starting at base of stack look for elements that require attention
	// base of stack is thought of as the start of the slice
	for i, v := range arry {
		if v == "-" {
			// Be smart about flipping by checking top layer before flip
			arry = w.setupTopOfStack(arry, i)

			ts := w.flip(arry[i:])
			bs := arry[:i]
			arry = w.appendTo(bs, ts)
		} else if v == "+" {
			// Do Nothing
		} else {
			// Handle errors on bad input
			errString := fmt.Sprintf("invalid '%s' at index: %v, only '-' or '+' allowed please correct input and retry\n", v, i)
			err = errors.New(errString)
			return 0, []string{}, err
		}
	}
	return w.flipCount, arry, err
}

// Start at top of stack and look for elements that require attention
// top of stack is thought of as the end of the slice
func (w *Waiter) setupTopOfStack(s []string, i int) []string {
	tc := 0
	for i := len(s) - 1; i >= i; i-- {
		if s[i] == "+" {
			tc++
		} else {
			break
		}
	}

	if tc > 0 {
		ts := w.flip(s[len(s)-tc:])
		bs := s[:len(s)-tc]
		s = w.appendTo(bs, ts)
	}
	return s
}

// flip swaps symbols and reverses slice
func (w *Waiter) flip(s []string) []string {
	// Update flipCount
	w.flipCount++

	// Invert strings
	for i, v := range s {
		if v == "-" {
			s[i] = "+"
		} else {
			s[i] = "-"
		}
	}
	s = w.reverse(s)
	return s
}

// reverse reverses the order of the elements in the slice
func (w *Waiter) reverse(s []string) []string {
	// Reverse slice
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// appendTo helps appends two slices together
func (w *Waiter) appendTo(s1 []string, s2 []string) []string {
	s := append(s1, s2...)
	return s
}
