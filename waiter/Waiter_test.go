package waiter_test

import (
	"math/rand"
	"testing"

	waiter "github.com/casonadams/go-pancakes/waiter"
)

var inputtest = []struct {
	input  []string
	output int
}{
	{[]string{"-"}, 1},
	{[]string{"-", "+"}, 2},
	{[]string{"+", "-"}, 1},
	{[]string{"+", "+", "+"}, 0},
	{[]string{"-", "-", "+", "-"}, 3},
}

func TestInputs(t *testing.T) {
	guy := waiter.NewWaiter()

	for i, tt := range inputtest {
		expected := tt.output
		actual, err := guy.Organize(tt.input)
		if err != nil {
			t.Errorf("Error: %v", err)
		} else {
			if actual != expected {
				t.Errorf("Test failed, expected: '%v', got: '%v', at iteration: %v", expected, actual, i)
			}
		}
	}
}

func BenchmarkWaiterFlip(b *testing.B) {
	options := []string{"-", "+"}
	slice := []string{options[rand.Intn(len(options))]}
	for i := 0; i < 10000; i++ {
		slice = append(slice, options[rand.Intn(len(options))])
	}

	guy := waiter.NewWaiter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		guy.Organize(slice)
	}
}
