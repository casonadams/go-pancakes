package waiter_test

import (
	"math/rand"
	"testing"

	waiter "github.com/casonadams/go-pancakes/waiter"
)

func getRandomData(s int) string {
	options := []string{"-", "+"}
	var stack string
	for i := 0; i < s; i++ {
		stack += options[rand.Intn(len(options))]
	}
	return stack
}

func TestInputs(t *testing.T) {
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

	for i, tt := range inputtest {
		expected := tt.output
		actual, _, err := w.Organize(tt.input)
		if err != nil {
			t.Errorf("Error: %v", err)
		} else {
			if actual != expected {
				t.Errorf("Test failed, expected: '%v', got: '%v', at iteration: %v", expected, actual, i)
			}
		}
	}
}

func TestLargeInput(t *testing.T) {
	slice := getRandomData(100)

	w := waiter.NewWaiter()
	_, stack, err := w.Organize(slice)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	for _, v := range stack {
		if v != "+" {
			t.Errorf("Test failed, expected to have all symbols b '+'")
		}
	}
}

func BenchmarkWaiterFlip(b *testing.B) {
	slice := getRandomData(100)

	w := waiter.NewWaiter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Organize(slice)
	}
}
