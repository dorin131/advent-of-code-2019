package main

import "testing"

func TestExecute(t *testing.T) {
	inputs := []struct {
		given    []int
		expected []int
	}{
		{
			[]int{99},
			[]int{99},
		}, {
			[]int{1, 0, 0, 0, 99},
			[]int{2, 0, 0, 0, 99},
		}, {
			[]int{2, 3, 0, 3, 99},
			[]int{2, 3, 0, 6, 99},
		}, {
			[]int{2, 4, 4, 5, 99, 0},
			[]int{2, 4, 4, 5, 99, 9801},
		}, {
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}
	for _, input := range inputs {
		t.Run("Execute", func(t *testing.T) {
			res, _ := Execute(input.given)
			for i := range input.expected {
				if input.expected[i] != input.given[i] {
					t.Errorf("Incorrect result\nExpected: %v\nReceived: %v\n", input.expected, res)
				}
			}
		})
	}
}
