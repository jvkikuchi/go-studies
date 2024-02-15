package main

import (
	"slices"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{10, 2}, []int{0, 9}, []int{0, 3})

	want := []int{12, 9, 3}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
