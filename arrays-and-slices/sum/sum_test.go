package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Should return 15", func(t *testing.T) {
		numbers := []int{1,2,3,4,5}

		got := Sum(numbers)
		want := 15

		assertGot(t, got, want, numbers[:])

	})
}

func assertGot(t testing.TB, got, want int, array []int) {
	t.Helper()
	
	if(got != want) {
		t.Errorf("got: %q / want: %q / given, %v", got, want, array)
	}
}