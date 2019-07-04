package waiter

import (
	"errors"
	"fmt"
)

type Waiter struct {
	flipCount  int
	finalStack []string
}

func NewWaiter() *Waiter {
	return &Waiter{}
}

func (w *Waiter) Organize(s []string) (int, error) {
	w.flipCount = 0
	var err error
	for i, v := range s {
		if v == "-" {
			// Be smart about flipping by checking top layer before flip
			tc := 0
			for k := len(s) - 1; k >= i; k-- {
				if s[k] == "+" {
					tc++
				} else {
					break
				}
			}
			if tc > 0 {
				ts := invertSliceAdj(s[len(s)-tc:])
				bs := s[:len(s)-tc]
				s = appendToAdj(bs, ts)
				w.flipCount++
			}
			ts := invertSliceAdj(s[i:])
			bs := s[:i]
			s = appendToAdj(bs, ts)
			w.flipCount++
		} else if v == "+" {
			// Do Nothing
		} else {
			// Need some error handeling on bad input
			errString := fmt.Sprintf("invalid '%s' at index: %v, only '-' or '+' allowed please correct input and retry\n", v, i)
			err = errors.New(errString)
			return 0, err
		}
	}
	return w.flipCount, err
}

func invertSliceAdj(s []string) []string {
	// Invert strings
	for i, v := range s {
		if v == "-" {
			s[i] = "+"
		} else {
			s[i] = "-"
		}
	}
	// Reverse slice
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func appendToAdj(s1 []string, s2 []string) []string {
	return append(s1, s2...)
}
