package repeat

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("it should repeat 3 times", func(t *testing.T) {
		got := Repeat("a", 2)
		want := "aa"
	
		assertGot(t, got, want)
	})

	t.Run("it should repeat 10 times", func(t *testing.T) {
		got := Repeat("1", 10)
		want := "1111111111"

		assertGot(t, got, want)
	})
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}

func assertGot(t testing.TB, got, want string) {
	t.Helper()
	
	if(got != want) {
		t.Errorf("got: %q / want: %q", got, want)
	}
}